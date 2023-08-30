package repository

import (
	"fmt"
	"log"
	"main/server/model"
	"main/server/pkg/dbclient"

	"github.com/jmoiron/sqlx"
)

type IUsersRepository interface {
	GetActiveSegments(userId int) (model.SegmentServiceModel, error)
	ChangeSegments() error
}

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *dbclient.Client) *UsersRepository {
	return &UsersRepository{
		db: db.Db,
	}
}

func (ur *UsersRepository) ChangeSegments(segments model.DbChangedSegments) error {
	var (
		isEmptySegTable bool
		err             error = nil
		isSegmentExists bool
	)
	query := `SELECT CheckSegmentsTableIsEmpty();`

	ur.db.Get(&isEmptySegTable, query)

	if isEmptySegTable {
		return fmt.Errorf("no seg data in table")
	}

	tx := ur.db.MustBegin()

	if !(len(segments.To_add) == 0) {
		query_exists := `SELECT CheckIsSegmentExists($1::text)`

		query = `CALL AddUserSegments($1::integer, $2::text, to_date((now()::date)::text, 'YYYY-MM-DD'));`

		for _, segment := range segments.To_add {
			ur.db.Get(&isSegmentExists, query_exists, segment)
			log.Println()
			if !isSegmentExists {
				log.Println("isSegmentExists error: ", isSegmentExists, "; segment: ", segment)
				err = fmt.Errorf("segment not exists")
			} else {
				_, err := tx.Exec(query, segments.User_id, segment)

				if err != nil {
					log.Println("Error while insert segment data; perhaps segment doesn't exist")
					err = fmt.Errorf("segment not exists")
					break
				}
			}

		}
	}

	if err != nil {
		return err
	}
	err = tx.Commit()

	if err != nil {
		return err
	}

	tx = ur.db.MustBegin()

	if !(len(segments.To_delete) == 0) {
		query := `CALL DeleteUserSegments($1::integer, $2::text, to_date((now()::date)::text, 'YYYY-MM-DD'));`

		for _, segment := range segments.To_delete {
			tx.Exec(query, segments.User_id, segment)
		}
	}
	err = tx.Commit()

	return err
}

func (ur *UsersRepository) GetActiveSegments(userId int) (*model.SegmentServiceModel, error) {
	var userSegment string
	var userSegments []string
	if userId <= 0 {
		return nil, fmt.Errorf("No user id")
	}

	tx := ur.db.MustBegin()
	query := `SELECT segment_name FROM users_segments as us WHERE us.user_id = $1 AND deleted_at IS NULL;`

	rows, err := tx.Queryx(query, userId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&userSegment)
		userSegments = append(userSegments, userSegment)
	}

	userSegmentsEntity := model.SegmentDbEntity{
		Segments: userSegments,
	}

	return model.SegEntityToModel(userSegmentsEntity), nil
}

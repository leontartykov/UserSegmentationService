package repository

import (
	"fmt"
	"log"
	"main/server/pkg/dbclient"
	"time"

	"github.com/jmoiron/sqlx"
)

type IUsersRepository interface {
	GetActiveSegments() error
	ChangeSegments() error
}

type UsersRepository struct {
	db *sqlx.DB
}

type DbSegments struct {
	to_add     []string
	to_delete  []string
	added_at   time.Time
	deleted_at time.Time
	user_id    string
}

func NewUsersRepository(db *dbclient.Client) *UsersRepository {
	return &UsersRepository{
		db: db.Db,
	}
}

func (ur *UsersRepository) ChangeSegments(segments DbSegments) error {
	if len(segments.to_add) == 0 || segments.added_at.IsZero() {
		return fmt.Errorf("no data to add segments in db segments")
	} else if len(segments.to_delete) == 0 || segments.deleted_at.IsZero() {
		return fmt.Errorf("no data to delete segments in db segments")
	}

	tx := ur.db.MustBegin()

	if len(segments.to_add) != 0 {
		query := `INSERT INTO usersSegments(userName, segmentName, added_at) VALUES ($1, $2, $3);`

		for _, segment := range segments.to_add {
			result, _ := tx.Exec(query, segments.user_id, segment, segments.added_at)
			log.Println(result.RowsAffected())
		}
	}

	if len(segments.to_delete) != 0 {
		query := `CALL changeUserSegments($1::integer, $2::text, to_date($3::text, 'YYYY-MM-DD'), to_date($4::text, 'YYYY-MM-DD')); `

		for _, segment := range segments.to_delete {
			result := tx.MustExec(query, segments.user_id, segment, segments.added_at, segments.deleted_at)
			log.Println(result)
			//tx.Rollback()
		}
	}
	err := tx.Commit()

	return err
}

func (ur *UsersRepository) GetActiveSegments() error {
	return nil
}

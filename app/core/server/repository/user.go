package repository

import (
	"log"
	"main/server/pkg/dbclient"

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
	to_add    []string
	to_delete []string
	user_id   string
}

func NewUsersRepository(db *dbclient.Client) *UsersRepository {
	return &UsersRepository{
		db: db.Db,
	}
}

func (ur *UsersRepository) ChangeSegments(segments DbSegments) error {
	tx := ur.db.MustBegin()

	if !(len(segments.to_add) == 0) {
		query := `CALL AddUserSegments($1::integer, $2::text, to_date((now()::date)::text, 'YYYY-MM-DD'));`

		for _, segment := range segments.to_add {
			tx.MustExec(query, segments.user_id, segment)
		}
	}

	if !(len(segments.to_delete) == 0) {
		query := `CALL DeleteUserSegments($1::integer, $2::text, to_date((now()::date)::text, 'YYYY-MM-DD'));`

		for _, segment := range segments.to_delete {
			tx.MustExec(query, segments.user_id, segment)
		}
	}
	err := tx.Commit()
	log.Println(err)

	return err
}

func (ur *UsersRepository) GetActiveSegments() error {
	return nil
}

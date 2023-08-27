package repository

import (
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

func NewUsersRepository(db *dbclient.Client) *UsersRepository {
	return &UsersRepository{
		db: db.Db,
	}
}

func GetActiveSegments() error {
	return nil
}

func ChangeSegments() error {
	return nil
}

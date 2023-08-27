package repository

import (
	"main/server/pkg/dbclient"

	"github.com/jmoiron/sqlx"
)

type ISegmentsRepository interface {
	Create(segment string) error
	Delete(segment string) error
}

type SegmentsRepository struct {
	db *sqlx.DB
}

func NewSegmentsRepository(db *dbclient.Client) *SegmentsRepository {
	return &SegmentsRepository{
		db: db.Db,
	}
}

func (sq *SegmentsRepository) Create(segment string) error {
	_, err := sq.db.NamedExec(`INSERT INTO segments (name) VALUES ($1)`, segment)
	return err
}

func (sq *SegmentsRepository) Delete(segment string) error {
	return nil
}

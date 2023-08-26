package repository

import "github.com/jmoiron/sqlx"

type SegmentsRepository struct {
	db *sqlx.DB
}

func NewSegmentRepository(db *sqlx.DB) *SegmentsRepository {
	return &SegmentsRepository{
		db: db,
	}
}

func (sq *SegmentsRepository) Create(segment string) error {
	_, err := sq.db.NamedExec(`INSERT INTO segments (name) VALUES ($1)`, segment)
	return err
}

func (sq *SegmentsRepository) Delete(segment string) error {
	return nil
}

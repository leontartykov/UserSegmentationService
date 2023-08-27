package repository

import (
	"fmt"
	"main/server/pkg/dbclient"

	"github.com/jmoiron/sqlx"
)

type ISegmentsRepository interface {
	Create(segmentName string) error
	Delete(segmentName string) error
}

type SegmentsRepository struct {
	db *sqlx.DB
}

func NewSegmentsRepository(db *dbclient.Client) *SegmentsRepository {
	return &SegmentsRepository{
		db: db.Db,
	}
}

func (sq *SegmentsRepository) Create(segmentName string) error {
	if segmentName == "" {
		return fmt.Errorf("failed while get empty segmentName")
	}

	_, err := sq.db.NamedExec(`INSERT INTO segments (name) VALUES (:name);`,
		map[string]interface{}{
			"name": segmentName,
		})
	return err
}

func (sq *SegmentsRepository) Delete(segmentName string) error {
	if segmentName == "" {
		return fmt.Errorf("failed while get empty segmentName")
	}

	_, err := sq.db.Exec(`DELETE FROM segments WHERE name =$1;`, segmentName)

	return err
}

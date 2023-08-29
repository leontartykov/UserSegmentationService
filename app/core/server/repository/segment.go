package repository

import (
	"fmt"
	"main/server/pkg/dbclient"
	"strings"

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

func (sr *SegmentsRepository) Create(segmentName string) error {
	if segmentName == "" {
		return fmt.Errorf("failed while get empty segmentName")
	}

	_, err := sr.db.NamedExec(`INSERT INTO segments (name) VALUES (:name);`,
		map[string]interface{}{
			"name": segmentName,
		})

	if strings.Contains(fmt.Sprint(err), `pq: duplicate key value violates unique constraint "c_segments_name_unique`) {
		err = fmt.Errorf("dublicate value")
	}

	return err
}

func (sr *SegmentsRepository) Delete(segmentName string) error {
	if segmentName == "" {
		return fmt.Errorf("failed while get empty segmentName")
	}

	_, err := sr.db.Exec(`DELETE FROM segments WHERE name =$1;`, segmentName)

	return err
}

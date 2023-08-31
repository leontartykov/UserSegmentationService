package repository

import (
	"fmt"
	"log"
	"main/server/model"
	"main/server/pkg/dbclient"
	"strings"

	"github.com/jmoiron/sqlx"
)

type ISegmentsRepository interface {
	Create(segmentName string) error
	Delete(segmentName string) error
	CreateWithUserPercent(segmentName string, usersId model.UsersWithNeedSegmentDb) error
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

func (sr *SegmentsRepository) CreateWithUserPercent(segmentName string, usersId model.UsersWithNeedSegmentDb) error {
	var (
		isSegmentExists bool
		err             error
	)

	if segmentName == "" {
		return fmt.Errorf("failed while get empty segmentName")
	}

	query_exists := `SELECT CheckIsSegmentExists($1::text)`
	query := `CALL AddUserSegments($1::integer, $2::text, to_date((now()::date)::text, 'YYYY-MM-DD'));`

	tx := sr.db.MustBegin()

	sr.db.Get(&isSegmentExists, query_exists, segmentName)

	if !isSegmentExists {
		log.Println("isSegmentExists error: ", isSegmentExists, "; segment: ", segmentName)
		err = fmt.Errorf("segment not exists")
	} else {
		for _, id := range usersId.UsersId {
			tx.Exec(query, id, segmentName)
		}
	}

	if err != nil {
		return err
	}
	err = tx.Commit()

	return err
}

package repository

import (
	"fmt"
	"log"
	"main/server/model"
	"main/server/pkg/dbclient"

	"github.com/jmoiron/sqlx"
)

type IReportsRepository interface {
	GetUsersWithSegments(dateTime string) ([]model.ReportEntityDb, error)
}

type ReportsRepository struct {
	db *sqlx.DB
}

func NewReportsRepository(db *dbclient.Client) *ReportsRepository {
	return &ReportsRepository{
		db: db.Db,
	}
}

func (rr *ReportsRepository) GetUsersWithSegments(dateTime string) ([]model.ReportEntityServ, error) {
	var (
		reportEntity model.ReportEntityDb
		report       []model.ReportEntityDb
	)

	if dateTime == "" {
		return nil, fmt.Errorf("Error: empty date time")
	}

	query := fmt.Sprintf("SELECT * FROM get_report_about_users_segments('%s');", dateTime)
	rows, err := rr.db.Queryx(query)

	if err != nil {
		log.Println("Problem with get report about users segments in repository")
		return nil, fmt.Errorf("Problem with get report about users segments")

	}
	for rows.Next() {
		err := rows.StructScan(&reportEntity)
		if err != nil {
			log.Println("Problem with get row in report about users segments in repository")
			return nil, fmt.Errorf("Problem with get report about users segments")
		}
		report = append(report, reportEntity)
	}

	return model.ReportEntityDbToServ(report), nil
}

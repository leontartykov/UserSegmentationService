package repository

import (
	"fmt"
	"log"
	"main/server/pkg/dbclient"
	"main/server/session"
	"testing"
)

func TestReportsRepositoryInterface(t *testing.T) {
	config := session.GetConfig()
	dbClient, err := dbclient.NewDbConnection(&config.DB)

	if err != nil {
		log.Fatal(err)
	}

	initDataBaseToReportTests(dbClient)
	reportsRepo := NewReportsRepository(dbClient)

	t.Run("GenerateReportAboutUsersAndSegments", func(t *testing.T) {
		date_report := "2023-08-02"
		_, err := reportsRepo.GetUsersWithSegments(date_report)

		if err != nil {
			t.Errorf(fmt.Sprint(err))
		}
	})
}

func initDataBaseToReportTests(dbClient *dbclient.Client) error {
	tx := dbClient.Db.MustBegin()

	query := `INSERT INTO users_segments VALUES 
				(1, 'segment_1', '2023-08-31', '2023-09-01'),
				(1, 'segment_2', '2023-07-04', null),
				(1, 'segment_3', '2023-08-01', '2023-09-06'),
				(2, 'segment_1', '2023-09-30', '2023-10-01'),
				(2, 'segment_2', '2023-07-25', '2023-08-01'),
				(2, 'segment_3', '2023-07-26', '2023-08-02');`

	tx.Exec(query)
	err := tx.Commit()

	return err
}

package repository

import (
	"log"
	"main/server/model"
	"main/server/pkg/dbclient"
	"main/server/session"
	"testing"
)

func TestUsersRepositoryInterface(t *testing.T) {
	config := session.GetConfig()
	dbClient, err := dbclient.NewDbConnection(&config.DB)

	if err != nil {
		log.Fatal(err)
	}

	userRepo := NewUsersRepository(dbClient)

	dbSegmentsAdd := model.DbChangedSegments{
		To_add:  []string{"AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS", "AVITO_IMAGES", "AVITO_DISCOUNT_30", "AVITO_DISCOUNT_50"},
		User_id: "1",
	}

	dbSegmentsDelete := model.DbChangedSegments{
		To_delete: []string{"AVITO_DISCOUNT_30", "AVITO_PERFORMANCE_VAS"},
		User_id:   "1",
	}

	dbSegmentsEmpty := model.DbChangedSegments{
		User_id: "1",
	}

	dbSegmentsAddSame := model.DbChangedSegments{
		To_add:  []string{"AVITO_VOICE_MESSAGES", "AVITO_IMAGES"},
		User_id: "1",
	}

	dbSegmentsAddSameAfterDelete := model.DbChangedSegments{
		To_add:  []string{"AVITO_DISCOUNT_30", "AVITO_PERFORMANCE_VAS"},
		User_id: "1",
	}

	err = initDataBase(dbClient, dbSegmentsAdd)
	if err != nil {
		log.Fatal(err)
	}

	segmentsRepo := NewSegmentsRepository(dbClient)
	for _, segment := range dbSegmentsAdd.To_add {
		segmentsRepo.Create(segment)
	}

	t.Run("Unit=ChangeUserSegmentsAdd", func(t *testing.T) {
		err := userRepo.ChangeSegments(dbSegmentsAdd)

		if err != nil {
			t.Error(err)
		}

		tx := dbClient.Db.MustBegin()
		query := `SELECT from users_segments`
		result := tx.MustExec(query)

		get_rows, _ := result.RowsAffected()

		if int(get_rows) != len(dbSegmentsAdd.To_add) {
			t.Errorf("Not right number of added rows: expected %d, want %d", int(get_rows), len(dbSegmentsAdd.To_add))
		}
	})

	t.Run("Unit=ChangeUserSegmentsDelete", func(t *testing.T) {
		err := userRepo.ChangeSegments(dbSegmentsDelete)

		if err != nil {
			t.Error(err)
		}

		tx := dbClient.Db.MustBegin()
		query := `SELECT from users_segments`
		result := tx.MustExec(query)

		get_rows, _ := result.RowsAffected()

		if int(get_rows) != len(dbSegmentsAdd.To_add) {
			t.Errorf("Not right number of rows after delete: expected %d, want %d", int(get_rows), len(dbSegmentsAdd.To_add))
		}
	})

	t.Run("Unit=ChangeUserSegmentsEmpty", func(t *testing.T) {
		err := userRepo.ChangeSegments(dbSegmentsEmpty)

		if err != nil {
			t.Error(err)
		}

		tx := dbClient.Db.MustBegin()
		query := `SELECT from users_segments`
		result := tx.MustExec(query)

		get_rows, _ := result.RowsAffected()

		if int(get_rows) != len(dbSegmentsAdd.To_add) {
			t.Errorf("Not right number of rows after empty add: expected %d, want %d", int(get_rows), len(dbSegmentsAdd.To_add))
		}
	})

	t.Run("Unit=ChangeUserSegmentsAddSameSegment", func(t *testing.T) {
		err := userRepo.ChangeSegments(dbSegmentsAddSame)

		if err != nil {
			t.Error(err)
		}

		tx := dbClient.Db.MustBegin()
		query := `SELECT from users_segments`
		result := tx.MustExec(query)

		get_rows, _ := result.RowsAffected()

		if int(get_rows) != len(dbSegmentsAdd.To_add) {
			t.Errorf("Not right number of rows after add: expected %d, want %d", int(get_rows), len(dbSegmentsAdd.To_add)+1)
		}
	})

	t.Run("Unit=ChangeUserSegmentsAddSameSegmentAfterDelete", func(t *testing.T) {
		err := userRepo.ChangeSegments(dbSegmentsAddSameAfterDelete)

		if err != nil {
			t.Error(err)
		}

		tx := dbClient.Db.MustBegin()
		query := `SELECT from users_segments`
		result := tx.MustExec(query)

		get_rows, _ := result.RowsAffected()

		if int(get_rows) != len(dbSegmentsAdd.To_add)+len(dbSegmentsAddSameAfterDelete.To_add) {
			t.Errorf("Not right number of rows add after delete: expected %d, want %d", int(get_rows), len(dbSegmentsAdd.To_add)+len(dbSegmentsAddSameAfterDelete.To_add))
		}
	})

	t.Run("Unit=IncorrectUserId", func(t *testing.T) {
		userId := 0
		_, err := userRepo.GetActiveSegments(userId)

		if err == nil {
			t.Error(err)
		}
	})
}

func initDataBase(dbClient *dbclient.Client, segments model.DbChangedSegments) error {
	tx := dbClient.Db.MustBegin()

	query := `DELETE FROM users_segments;
			  DELETE FROM segments;`

	tx.MustExec(query)

	segmentsRepo := NewSegmentsRepository(dbClient)
	for _, segment := range segments.To_add {
		segmentsRepo.Create(segment)
	}

	err := tx.Commit()

	return err
}

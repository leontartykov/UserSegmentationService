package repository

import (
	"log"
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

	dbSegmentsAdd := DbSegments{
		to_add:  []string{"AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS", "AVITO_IMAGES", "AVITO_DISCOUNT_30", "AVITO_DISCOUNT_50"},
		user_id: "1",
	}

	dbSegmentsDelete := DbSegments{
		to_delete: []string{"AVITO_DISCOUNT_30", "AVITO_PERFORMANCE_VAS"},
		user_id:   "1",
	}

	dbSegmentsEmpty := DbSegments{
		user_id: "1",
	}

	dbSegmentsAddSame := DbSegments{
		to_add:  []string{"AVITO_VOICE_MESSAGES", "AVITO_IMAGES"},
		user_id: "1",
	}

	dbSegmentsAddSameAfterDelete := DbSegments{
		to_add:  []string{"AVITO_DISCOUNT_30", "AVITO_PERFORMANCE_VAS"},
		user_id: "1",
	}

	t.Run("ChangeUserSegmentsAdd", func(t *testing.T) {
		err := userRepo.ChangeSegments(dbSegmentsAdd)

		if err != nil {
			t.Error(err)
		}

		tx := dbClient.Db.MustBegin()
		query := `SELECT from userssegments`
		result := tx.MustExec(query)

		get_rows, _ := result.RowsAffected()

		if int(get_rows) != len(dbSegmentsAdd.to_add) {
			t.Errorf("Not right number of added rows: expected %d, want %d", int(get_rows), len(dbSegmentsAdd.to_add))
		}
	})

	t.Run("ChangeUserSegmentsDelete", func(t *testing.T) {
		err := userRepo.ChangeSegments(dbSegmentsDelete)

		if err != nil {
			t.Error(err)
		}

		tx := dbClient.Db.MustBegin()
		query := `SELECT from userssegments`
		result := tx.MustExec(query)

		get_rows, _ := result.RowsAffected()

		if int(get_rows) != len(dbSegmentsAdd.to_add) {
			t.Errorf("Not right number of rows after delete: expected %d, want %d", int(get_rows), len(dbSegmentsAdd.to_add))
		}
	})

	t.Run("ChangeUserSegmentsEmpty", func(t *testing.T) {
		err := userRepo.ChangeSegments(dbSegmentsEmpty)

		if err != nil {
			t.Error(err)
		}

		tx := dbClient.Db.MustBegin()
		query := `SELECT from userssegments`
		result := tx.MustExec(query)

		get_rows, _ := result.RowsAffected()

		if int(get_rows) != len(dbSegmentsAdd.to_add) {
			t.Errorf("Not right number of rows after empty add: expected %d, want %d", int(get_rows), len(dbSegmentsAdd.to_add))
		}
	})

	t.Run("ChangeUserSegmentsAddSameSegment", func(t *testing.T) {
		err := userRepo.ChangeSegments(dbSegmentsAddSame)

		if err != nil {
			t.Error(err)
		}

		tx := dbClient.Db.MustBegin()
		query := `SELECT from userssegments`
		result := tx.MustExec(query)

		get_rows, _ := result.RowsAffected()

		if int(get_rows) != len(dbSegmentsAdd.to_add) {
			t.Errorf("Not right number of rows after add: expected %d, want %d", int(get_rows), len(dbSegmentsAdd.to_add)+1)
		}
	})

	t.Run("ChangeUserSegmentsAddSameSegmentAfterDelete", func(t *testing.T) {
		err := userRepo.ChangeSegments(dbSegmentsAddSameAfterDelete)

		if err != nil {
			t.Error(err)
		}

		tx := dbClient.Db.MustBegin()
		query := `SELECT from userssegments`
		result := tx.MustExec(query)

		get_rows, _ := result.RowsAffected()

		if int(get_rows) != len(dbSegmentsAdd.to_add)+len(dbSegmentsAddSameAfterDelete.to_add) {
			t.Errorf("Not right number of rows add after delete: expected %d, want %d", int(get_rows), len(dbSegmentsAdd.to_add)+len(dbSegmentsAddSameAfterDelete.to_add))
		}
	})

	/*t.Run("ChangeUserSegmentsDelete", func(t *testing.T) {
		err := userRepo.ChangeSegments(dbSegmentsDelete)

		if err != nil {
			t.Errorf("No segments; want 1 more")
		}
	})*/
}

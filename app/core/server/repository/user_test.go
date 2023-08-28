package repository

import (
	"log"
	"main/server/pkg/dbclient"
	"main/server/session"
	"testing"
	"time"
)

func TestUsersRepositoryInterface(t *testing.T) {
	config := session.GetConfig()
	dbClient, err := dbclient.NewDbConnection(&config.DB)

	if err != nil {
		log.Fatal(err)
	}

	userRepo := NewUsersRepository(dbClient)

	dbSegmentsAdd := DbSegments{
		to_add:     []string{"AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS", "AVITO_IMAGES"},
		to_delete:  []string{"AVITO_DISCOUNT_30", "AVITO_DISCOUNT_50"},
		added_at:   time.Now(),
		deleted_at: time.Now(),
		user_id:    "1",
	}

	/*dbSegmentsDelete := DbSegments{
		to_delete:  []string{"AVITO_DISCOUNT_30", "AVITO_DISCOUNT_50"},
		deleted_at: time.Now(),
		user_id:    "1",
	}*/

	log.Println(dbSegmentsAdd)

	t.Run("ChangeUserSegmentsAdd", func(t *testing.T) {
		err := userRepo.ChangeSegments(dbSegmentsAdd)

		if err != nil {
			t.Error(err)
		}
	})

	/*t.Run("ChangeUserSegmentsDelete", func(t *testing.T) {
		err := userRepo.ChangeSegments(dbSegmentsDelete)

		if err != nil {
			t.Errorf("No segments; want 1 more")
		}
	})*/
}

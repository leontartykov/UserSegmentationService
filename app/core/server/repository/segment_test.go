package repository

import (
	"log"
	"main/server/pkg/dbclient"
	"main/server/session"
	"testing"
)

func TestSegmentsRepositoryInterface(t *testing.T) {
	config := session.GetConfig()
	dbClient, err := dbclient.NewDbConnection(&config.DB)

	if err != nil {
		log.Fatal(err)
	}

	segmentsRepo := NewSegmentsRepository(dbClient)

	exampleSegName := "AVITO_VOICE_MESSAGES"

	t.Run("CreateSegmentWithName", func(t *testing.T) {
		err := segmentsRepo.Create(exampleSegName)

		if err != nil {
			t.Errorf("No segment name; want 1")
		}
	})

	t.Run("CreateSegmentNoName", func(t *testing.T) {
		err := segmentsRepo.Create("")

		if err == nil {
			t.Errorf("Is segment name; want 0")
		}
	})

	t.Run("DeleteSegment", func(t *testing.T) {
		err := segmentsRepo.Delete(exampleSegName)

		if err != nil {
			t.Errorf("No segment name to delete; want 1")
		}
	})
}

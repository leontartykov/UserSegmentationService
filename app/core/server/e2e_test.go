package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"main/server/handlers"
	"main/server/pkg/dbclient"
	"main/server/repository"
	"main/server/services"
	"main/server/session"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type resultJsonMessage struct {
	status string
}

var (
	createSegmentRegisterBody = `{"name": "AVITO_VOICE_MESSAGES"}`
	result                    resultJsonMessage
	changeSegs                = `{"to_add": ["AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS", "AVITO_DISCOUNT_30"]}`
)

func TestApiV1Methods(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	config := session.GetConfig()
	dbClient, err := dbclient.NewDbConnection(&config.DB)

	if err != nil {
		log.Fatal(err)
	}
	segmentsRepo := repository.NewSegmentsRepository(dbClient)
	segmentsServ := services.NewSegmentsService(*segmentsRepo)
	segmentsHandler := handlers.NewSegmentsHandler(*segmentsServ)

	usersRepo := repository.NewUsersRepository(dbClient)
	usersServ := services.NewUsersService(*usersRepo)
	usersHandler := handlers.NewUsersHandler(*usersServ)

	segmentsHandler.Register(router)
	usersHandler.Register(router)

	clearDataBase(dbClient)

	t.Run("E2E: create segment", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/api/v1/segments", bytes.NewBufferString(createSegmentRegisterBody))

		router.ServeHTTP(w, r)

		assert.Equal(t, 201, w.Code)

		json.NewDecoder(w.Body).Decode(&result)
		log.Println("result: ", result)
	})

	t.Run("E2E: delete same segment", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/segments/%s", "AVITO_VOICE_MESSAGES"), bytes.NewBufferString(createSegmentRegisterBody))

		router.ServeHTTP(w, r)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("E2E: create segment that already exists", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/api/v1/segments", bytes.NewBufferString(createSegmentRegisterBody))

		router.ServeHTTP(w, r)

		w = httptest.NewRecorder()
		r = httptest.NewRequest("POST", "/api/v1/segments", bytes.NewBufferString(createSegmentRegisterBody))

		router.ServeHTTP(w, r)

		assert.Equal(t, 201, w.Code)

		json.NewDecoder(w.Body).Decode(&result)
		log.Println("result: ", result)
	})

	t.Run("E2E: add not exists segments to user", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("PUT", "/api/v1/users/1/segments", bytes.NewBufferString(changeSegs))

		router.ServeHTTP(w, r)

		assert.Equal(t, 404, w.Code)
	})

	t.Run("E2E: add exists segments to user", func(t *testing.T) {
		err := insertSomeSegments(dbClient)
		if err != nil {
			t.Errorf(fmt.Sprintf("Problem with insert data: %s", fmt.Sprint(err)))
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest("PUT", "/api/v1/users/1/segments", bytes.NewBufferString(changeSegs))

		router.ServeHTTP(w, r)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("E2E: get user's active segments", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/api/v1/users/1/segments/active", nil)

		router.ServeHTTP(w, r)

		assert.Equal(t, w.Code, 200)
	})
}

func clearDataBase(dbClient *dbclient.Client) error {
	tx := dbClient.Db.MustBegin()

	query := `DELETE FROM users_segments;
			  DELETE FROM segments;`

	tx.MustExec(query)

	err := tx.Commit()

	return err
}

func insertSomeSegments(dbClient *dbclient.Client) error {
	tx := dbClient.Db.MustBegin()
	query := `DELETE FROM segments;
			  INSERT INTO segments (name) VALUES ('AVITO_VOICE_MESSAGES'), ('AVITO_PERFORMANCE_VAS'), ('AVITO_DISCOUNT_30');`
	tx.MustExec(query)
	err := tx.Commit()

	return err
}

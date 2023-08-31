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
	Status string
}

type resultJsonActiveSegments struct {
	Segments []string
}

var (
	createSegmentRegisterBody = `{"name": "AVITO_VOICE_MESSAGES"}`
	resultMsg                 resultJsonMessage
	changeSegsToAdd           = `{"to_add": ["AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS", "AVITO_DISCOUNT_30"]}`
	changeSegsToDelete        = `{"to_delete": ["AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS", "AVITO_DISCOUNT_30"]}`
	resultActiveSegs          resultJsonActiveSegments
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

	reportsRepo := repository.NewReportsRepository(dbClient)
	reportsServ := services.NewReportsService(*reportsRepo)
	reportsHandler := handlers.NewReportsHandler(*reportsServ)

	segmentsHandler.Register(router)
	usersHandler.Register(router)
	reportsHandler.Register(router)

	clearDataBase(dbClient)

	t.Run("E2E: create segment", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/api/v1/segments", bytes.NewBufferString(createSegmentRegisterBody))

		router.ServeHTTP(w, r)

		assert.Equal(t, 201, w.Code)

		json.NewDecoder(w.Body).Decode(&resultMsg)
		log.Println("result: ", resultMsg)
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

		json.NewDecoder(w.Body).Decode(&resultMsg)
		log.Println("result: ", resultMsg)
	})

	t.Run("E2E: add not exists segments to user", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("PUT", "/api/v1/users/1/segments", bytes.NewBufferString(changeSegsToAdd))

		router.ServeHTTP(w, r)

		assert.Equal(t, 404, w.Code)
	})

	t.Run("E2E: add exists segments to user", func(t *testing.T) {
		err := insertSomeSegments(dbClient)
		if err != nil {
			t.Errorf(fmt.Sprintf("Problem with insert data: %s", fmt.Sprint(err)))
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest("PUT", "/api/v1/users/1/segments", bytes.NewBufferString(changeSegsToAdd))

		router.ServeHTTP(w, r)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("E2E: get user's active segments", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/api/v1/users/1/segments/active", nil)

		router.ServeHTTP(w, r)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("E2E: delete user's segments", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("PUT", "/api/v1/users/1/segments", bytes.NewBufferString(changeSegsToDelete))

		router.ServeHTTP(w, r)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("E2E: get user's active segments after delete user from them all", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/api/v1/users/1/segments/active", nil)

		router.ServeHTTP(w, r)

		assert.Equal(t, 200, w.Code)

		json.NewDecoder(w.Body).Decode(&resultActiveSegs)

		assert.Equal(t, 0, len(resultActiveSegs.Segments))
	})

	t.Run("E2E: add exists segments to user", func(t *testing.T) {
		err := insertSomeSegments(dbClient)
		if err != nil {
			t.Errorf(fmt.Sprintf("Problem with insert data: %s", fmt.Sprint(err)))
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest("PUT", "/api/v1/users/1/segments", bytes.NewBufferString(changeSegsToAdd))

		router.ServeHTTP(w, r)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("E2E: get user's active segments after deleting and inserting same", func(t *testing.T) {
		needActiveSegs := []string{"AVITO_VOICE_MESSAGES", "AVITO_PERFORMANCE_VAS", "AVITO_DISCOUNT_30"}
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/api/v1/users/1/segments/active", nil)

		router.ServeHTTP(w, r)

		assert.Equal(t, 200, w.Code)

		json.NewDecoder(w.Body).Decode(&resultActiveSegs)

		assert.Equal(t, len(needActiveSegs), len(resultActiveSegs.Segments))

		for i, segment := range resultActiveSegs.Segments {
			assert.Equal(t, needActiveSegs[i], segment)
		}
	})

	t.Run("E2E: get report about users and segments", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/api/v1/reports/usersSegs/2023-08", nil)

		router.ServeHTTP(w, r)
		assert.Equal(t, 200, w.Code)
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

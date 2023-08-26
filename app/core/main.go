package main

import (
	"log"
	"main/server"
	"main/server/handlers"
	"main/server/pkg/dbclient"
	"main/server/session"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	v1 := router.Group("/api/v1")
	{
		segments := v1.Group("/segments")
		{
			segments.POST("", handlers.CreateSegment)
			segments.DELETE(":name", handlers.DeleteSegment)
		}

		users := v1.Group("/users/:id")
		{
			users.PUT("/segments", handlers.ChangeUserSegments)
			users.GET("/segments/active", handlers.GetActiveUserSegments)
		}
	}

	config := session.GetConfig()
	_, err := dbclient.NewDbConnection(&config.DB)

	if err != nil {
		log.Fatal(err)
	}

	server.Run(config, router)

}

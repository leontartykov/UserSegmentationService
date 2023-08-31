package main

import (
	"log"
	"main/server"
	"main/server/handlers"
	"main/server/pkg/dbclient"
	"main/server/repository"
	"main/server/services"
	"main/server/session"

	"github.com/gin-gonic/gin"
)

func main() {
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

	server.Run(config, router)

}

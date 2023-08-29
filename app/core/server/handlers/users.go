package handlers

import (
	"log"
	"main/server/model"
	"main/server/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IUsersHandler interface {
	ChangeUserSegments(c *gin.Context)
	GetActiveUserSegments(c *gin.Context)
}

type UsersHandler struct {
	service services.UsersService
}

func NewUsersHandler(service services.UsersService) *UsersHandler {
	return &UsersHandler{
		service: service,
	}
}

func (uh *UsersHandler) Register(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users/:id")
		{
			users.PUT("/segments", uh.ChangeUserSegments)
			users.GET("/segments/active", uh.GetActiveUserSegments)
		}
	}
}

func (uh *UsersHandler) ChangeUserSegments(c *gin.Context) {
	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "no input id data"})
		return
	}

	var requestBody model.ServChangedSegments

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "no input segments data"})
		return
	}

	requestBody.User_id = id
	log.Println(requestBody)

	err := uh.service.ChangeSegments(requestBody)

	log.Println("ERROR", err)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"status": "successful changing"})
		return
	} else {
		c.JSON(http.StatusBadGateway, gin.H{"status": "ooops"})
		return
	}
}

func (uh *UsersHandler) GetActiveUserSegments(c *gin.Context) {

}

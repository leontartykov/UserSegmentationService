package handlers

import (
	"fmt"
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

	err := uh.service.ChangeSegments(requestBody)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"status": "successful changing"})
	} else if fmt.Sprint(err) == "segment not exists" || fmt.Sprint(err) == "no seg data in table" || fmt.Sprint(err) == "one of segment to delete not found" {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"status": "one of segments not exists"})
	} else {
		log.Println("Error: ", fmt.Sprint(err))
		c.JSON(http.StatusBadGateway, gin.H{"status": "ooops"})
	}
	return
}

func (uh *UsersHandler) GetActiveUserSegments(c *gin.Context) {
	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "no input id data"})
		return
	}

	segments, err := uh.service.GetActiveSegments(id)

	if err == nil {
		c.JSON(http.StatusOK, segments)
	} else {
		log.Println("Error: ", fmt.Sprint(err))
		c.JSON(http.StatusBadGateway, gin.H{"status": "ooops"})
	}

	return
}

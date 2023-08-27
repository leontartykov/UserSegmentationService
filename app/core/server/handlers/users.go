package handlers

import (
	"main/server/services"

	"github.com/gin-gonic/gin"
)

type IUsersHandler interface {
	ChangeUserSegments(context *gin.Context)
	GetActiveUserSegments(context *gin.Context)
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
			users.PUT("/segments", uh.service.ChangeSegments)
			users.GET("/segments/active", uh.service.GetActiveSegments)
		}
	}
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandler interface {
	ChangeUserSegments(context *gin.Context)
	GetActiveUserSegments(context *gin.Context)
}

func ChangeUserSegments(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "put"})
}

func GetActiveUserSegments(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "get"})
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SegmentsHandler interface {
	CreateSegment(context *gin.Context)
	DeleteSegment(context *gin.Context)
}

func CreateSegment(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func DeleteSegment(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{"status": "deleted"})
}

package services

import (
	"main/server/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ISegmentsService interface {
	CreateSegment(context *gin.Context)
	DeleteSegment(context *gin.Context)
}

type SegmentsService struct {
	repository repository.SegmentsRepository
}

func NewSegmentsService(repository repository.SegmentsRepository) *SegmentsService {
	return &SegmentsService{
		repository: repository,
	}
}

func (ss *SegmentsService) CreateSegment(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (ss *SegmentsService) DeleteSegment(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{"status": "deleted"})
}

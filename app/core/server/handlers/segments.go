package handlers

import (
	"main/server/services"

	"github.com/gin-gonic/gin"
)

type SegmentsHandler struct {
	service services.SegmentsService
}

func NewSegmentsHandler(service services.SegmentsService) *SegmentsHandler {
	return &SegmentsHandler{
		service: service,
	}
}

func (sh *SegmentsHandler) Register(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		segments := v1.Group("/segments")
		{
			segments.POST("", sh.service.CreateSegment)
			segments.DELETE(":name", sh.service.DeleteSegment)
		}
	}
}

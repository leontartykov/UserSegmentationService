package handlers

import (
	"main/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ISegmentHandler interface {
	CreateSegment(c *gin.Context)
}

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
			segments.POST("", sh.CreateSegment)
			segments.DELETE(":name", sh.DeleteSegment)
		}
	}
}

type SegmentsRequestBody struct {
	SegmentName string
}

func (sh *SegmentsHandler) CreateSegment(c *gin.Context) {
	var requestBody SegmentsRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "no input data"})
		return
	}

	sh.service.CreateSegment(requestBody.SegmentName)
}

func (sh *SegmentsHandler) DeleteSegment(c *gin.Context) {

}

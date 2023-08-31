package handlers

import (
	"fmt"
	"log"
	"main/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ISegmentHandler interface {
	CreateSegment(c *gin.Context)
	DeleteSegment(c *gin.Context)
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
	Name string
}

func (sh *SegmentsHandler) CreateSegment(c *gin.Context) {
	var segRequestBody SegmentsRequestBody
	if err := c.BindJSON(&segRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "no input data"})
		return
	}

	err := sh.service.CreateSegment(segRequestBody.Name)

	log.Println(err)

	if err == nil {
		c.JSON(http.StatusCreated, gin.H{"status": "successful created"})
	} else if fmt.Sprint(err) == "failed to get segmentName" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed to get segment name"})
	} else if fmt.Sprint(err) == "dublicate value" {
		c.JSON(http.StatusCreated, gin.H{"status": "Ok"})
	} else {
		c.JSON(http.StatusBadGateway, gin.H{"status": "ooops"})
	}
	return
}

func (sh *SegmentsHandler) DeleteSegment(c *gin.Context) {
	segToDel := c.Param("name")
	err := sh.service.DeleteSegment(segToDel)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"status": "successful deleted"})
	} else if fmt.Sprint(err) == "failed to get segmentName" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed to get segment name"})
	} else {
		c.JSON(http.StatusBadGateway, gin.H{"status": "ooops"})
	}

	return
}

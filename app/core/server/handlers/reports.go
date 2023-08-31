package handlers

import (
	"encoding/csv"
	"fmt"
	"log"
	"main/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IReportsHandler interface {
	GetReportAboutUsersWithSegments(c *gin.Context)
}

type ReportsHandler struct {
	service services.ReportsService
}

func NewReportsHandler(service services.ReportsService) *ReportsHandler {
	return &ReportsHandler{
		service: service,
	}
}

func (sh *ReportsHandler) Register(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		segments := v1.Group("/reports/usersSegs/:year")
		{
			segments.GET("", sh.GetReportAboutUsersWithSegments)
		}
	}
}

func (sh *ReportsHandler) GetReportAboutUsersWithSegments(c *gin.Context) {
	yearReport := c.Param("year")

	report, err := sh.service.GetReportUsersWithSegments(yearReport)

	if err == nil {
		c.Header("Content-Type", "text/csv")
		c.Header("Content-Disposition", "attachment;filename=report.csv")

		wr := csv.NewWriter(c.Writer)

		if err := wr.WriteAll(report); err != nil {
			log.Println("Failed to generate CSV file")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate CSV file"})
			return
		} else {
			c.Status(http.StatusOK)
		}
	} else if fmt.Sprint(err) == "empty date time" {
		log.Println("Failed to get date time")
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed to get date time"})
	} else {
		log.Println(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": "ooops"})
	}
}

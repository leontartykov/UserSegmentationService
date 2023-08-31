package services

import (
	"fmt"
	"main/server/model"
	"main/server/repository"
)

type IReportsService interface {
	GetReportUsersWithSegments(dateTime string) error
}

type ReportsService struct {
	repository repository.ReportsRepository
}

func NewReportsService(repository repository.ReportsRepository) *ReportsService {
	return &ReportsService{
		repository: repository,
	}
}

func (rs *ReportsService) GetReportUsersWithSegments(dateTime string) ([][]string, error) {
	if dateTime == "" {
		return nil, fmt.Errorf("empty date time")
	}

	dateTimeServ := dateTime + "-01"
	report, err := rs.repository.GetUsersWithSegments(dateTimeServ)

	return model.ReportServToHandler(report), err
}

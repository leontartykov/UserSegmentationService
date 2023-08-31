package services

import (
	"fmt"
	"main/server/model"
	"main/server/repository"
)

type ISegmentsService interface {
	CreateSegment(segmentName string) error
	DeleteSegment(segmentName string) error
}

type SegmentsService struct {
	repository repository.SegmentsRepository
}

func NewSegmentsService(repository repository.SegmentsRepository) *SegmentsService {
	return &SegmentsService{
		repository: repository,
	}
}

func (ss *SegmentsService) CreateSegment(segment model.SegmentsCreateRequestBody) error {
	var err error
	if segment.Name == "" || (segment.Percent < 0 || segment.Percent > 100) {
		return fmt.Errorf("failed to get segmentName")
	}

	//TODO: mediator between segment Service and Users service
	if segment.Percent == 0 {
		err = ss.repository.Create(segment.Name)
	}

	return err
}

func (ss *SegmentsService) DeleteSegment(segmentName string) error {
	if segmentName == "" {
		return fmt.Errorf("failed to get segmentName")
	}

	return ss.repository.Delete(segmentName)
}

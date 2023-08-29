package services

import (
	"fmt"
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

func (ss *SegmentsService) CreateSegment(segmentName string) error {
	if segmentName == "" {
		return fmt.Errorf("failed to get segmentName")
	}

	return ss.repository.Create(segmentName)
}

func (ss *SegmentsService) DeleteSegment(segmentName string) error {
	if segmentName == "" {
		return fmt.Errorf("failed to get segmentName")
	}

	return ss.repository.Delete(segmentName)
}

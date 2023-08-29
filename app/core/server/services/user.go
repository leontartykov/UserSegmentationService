package services

import (
	"fmt"
	"log"
	"main/server/model"
	"main/server/repository"
	"strconv"
)

type IUserssService interface {
	ChangeSegments(segs model.ServChangedSegments) error
	GetActiveSegments(userId int) (*model.SegmentServiceModel, error)
}

type UsersService struct {
	repository repository.UsersRepository
}

func NewUsersService(repository repository.UsersRepository) *UsersService {
	return &UsersService{
		repository: repository,
	}
}

func (us *UsersService) ChangeSegments(segs model.ServChangedSegments) error {
	log.Println("Service struct: ", segs)
	if len(segs.To_add) == 0 && len(segs.To_delete) == 0 {
		return fmt.Errorf("Empty struct ")
	}

	err := us.repository.ChangeSegments(*model.ChangeSegsModelToEntity(segs))

	return err
}

func (us *UsersService) GetActiveSegments(userId string) (*model.SegmentServiceModel, error) {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}

	active_segs, err := us.repository.GetActiveSegments(id)

	return active_segs, err
}

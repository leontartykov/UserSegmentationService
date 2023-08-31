package services

import (
	"fmt"
	"log"
	"main/server/model"
	"main/server/repository"
	"math/rand"
	"strconv"
	"time"
)

type IUserssService interface {
	ChangeSegments(segs model.ServChangedSegments) error
	GetActiveSegments(userId int) (*model.SegmentServiceModel, error)
	GetUsersWithoutSegment(segmentName string, usersPercent int) ([]int, error)
}

type UsersService struct {
	repository repository.UsersRepository
	randGen    *rand.Rand
}

func NewUsersService(repository repository.UsersRepository) *UsersService {
	return &UsersService{
		repository: repository,
		randGen:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (us *UsersService) ChangeSegments(segs model.ServChangedSegments) error {
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

func (us *UsersService) GetUsersWithoutSegment(segmentName string, usersPercent int) ([]int, error) {
	var luckyPerson int
	if segmentName == "" {
		return nil, fmt.Errorf("no segment name")
	}

	ids, err := us.repository.GetUsersWithoutSegment(segmentName)

	if err != nil {
		return nil, err
	}

	nPersons := int(len(ids.UsersId) * (usersPercent / 100))

	luckyPeople := make([]int, nPersons)

	min, max := 0, len(ids.UsersId)-1

	for i := 0; i < nPersons; i++ {
		luckyPerson = us.randGen.Intn(max-min+1) + min
		luckyPeople[i] = ids.UsersId[luckyPerson]

		remove(ids.UsersId, luckyPerson)
		max -= 1
	}

	log.Println("lucky people: ", luckyPeople)
	log.Println("ids.UsersId: ", ids.UsersId)

	return luckyPeople, nil
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

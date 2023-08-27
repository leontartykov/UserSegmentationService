package services

import (
	"main/server/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserssService interface {
	ChangeSegments(context *gin.Context)
	GetActiveSegments(context *gin.Context)
}

type UsersService struct {
	repository repository.UsersRepository
}

func NewUsersService(repository repository.UsersRepository) *UsersService {
	return &UsersService{
		repository: repository,
	}
}

func (us *UsersService) ChangeSegments(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{"status": "put"})
}

func (us *UsersService) GetActiveSegments(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{"status": "get"})
}

package service

import (
	"github.com/Thunderbirrd/CourseProject/internal/repository"
	"github.com/Thunderbirrd/CourseProject/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Api interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id int) (models.User, error)
	GetUsersByServiceType(serviceType string) ([]models.User, error)
	GetUsersByLocation(location string) ([]models.User, error)
}

type User interface {
	UpdateUser(id int, input models.UpdateUserInput) error
}

type Service struct {
	Authorization
	Api
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Api:           NewApiService(repos.Api),
		User:          NewUserService(repos.User),
	}
}

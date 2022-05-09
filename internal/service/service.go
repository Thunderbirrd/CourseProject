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

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

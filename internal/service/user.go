package service

import (
	"github.com/Thunderbirrd/CourseProject/internal/repository"
	"github.com/Thunderbirrd/CourseProject/pkg/models"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) UpdateUser(id int, input models.UpdateUserInput) error {
	return s.repo.UpdateUser(id, input)
}

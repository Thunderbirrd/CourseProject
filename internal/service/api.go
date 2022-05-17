package service

import (
	"github.com/Thunderbirrd/CourseProject/internal/repository"
	"github.com/Thunderbirrd/CourseProject/pkg/models"
)

type ApiService struct {
	repo repository.Api
}

func NewApiService(repo repository.Api) *ApiService {
	return &ApiService{repo: repo}
}

func (s *ApiService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *ApiService) GetUserById(id int) (models.User, error) {
	return s.repo.GetUserById(id)
}

func (s *ApiService) GetUsersByServiceType(serviceType string) ([]models.User, error) {
	return s.repo.GetUsersByServiceType(serviceType)
}

func (s *ApiService) GetUsersByLocation(location string) ([]models.User, error) {
	return s.repo.GetUsersByLocation(location)
}

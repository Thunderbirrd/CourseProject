package repository

import (
	"github.com/Thunderbirrd/CourseProject/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Api interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id int) (models.User, error)
	GetUsersByServiceType(serviceType string) ([]models.User, error)
}

type User interface {
	UpdateUser(id int, input models.UpdateUserInput) error
}

type Repository struct {
	Authorization
	Api
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Api:           NewApiPostgres(db),
		User:          NewUserPostgres(db),
	}
}

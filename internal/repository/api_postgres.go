package repository

import (
	"fmt"
	"github.com/Thunderbirrd/CourseProject/pkg/models"
	"github.com/jmoiron/sqlx"
)

type ApiPostgres struct {
	db *sqlx.DB
}

func NewApiPostgres(db *sqlx.DB) *ApiPostgres {
	return &ApiPostgres{db: db}
}

func (r *ApiPostgres) GetAllUsers() ([]models.User, error) {
	var users []models.User

	query := fmt.Sprintf("SELECT id, name, username, service_type, info, location FROM %s", usersTable)

	if err := r.db.Select(&users, query); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *ApiPostgres) GetUserById(id int) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id, name, username, service_type, info, location FROM %s WHERE id = $1", usersTable)

	if err := r.db.Get(&user, query, id); err != nil {
		return user, err
	}

	return user, nil
}

func (r *ApiPostgres) GetUsersByServiceType(serviceType string) ([]models.User, error) {
	var users []models.User

	query := fmt.Sprintf("SELECT id, name, username, service_type, info, location FROM %s WHERE service_type = $1", usersTable)

	if err := r.db.Select(&users, query, serviceType); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *ApiPostgres) GetUsersByLocation(location string) ([]models.User, error) {
	var users []models.User

	query := fmt.Sprintf("SELECT id, name, username, service_type, info, location FROM %s WHERE location = $1", usersTable)

	if err := r.db.Select(&users, query, location); err != nil {
		return nil, err
	}

	return users, nil
}

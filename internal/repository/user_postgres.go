package repository

import (
	"fmt"
	"github.com/Thunderbirrd/CourseProject/pkg/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) UpdateUser(id int, input models.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Username != nil {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argId))
		args = append(args, *input.Username)
		argId++
	}

	if input.ServiceType != nil {
		setValues = append(setValues, fmt.Sprintf("service_type=$%d", argId))
		args = append(args, *input.ServiceType)
		argId++
	}

	if input.Info != nil {
		setValues = append(setValues, fmt.Sprintf("info=$%d", argId))
		args = append(args, *input.Info)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d",
		usersTable, setQuery, argId)

	args = append(args, id)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

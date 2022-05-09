package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

const usersTable = "users"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		log.Fatalf("Error while connecting to db: %s", err.Error())
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error while connecting to db: %s", err.Error())
		return nil, err
	}
	return db, nil
}

package models

type User struct {
	Id           int    `json:"-" db:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
	ServiceType  string `json:"service_type"`
	Info         string `json:"info"`
}

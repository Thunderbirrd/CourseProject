package models

type User struct {
	Id           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Username     string `json:"username" db:"username"`
	PasswordHash string `json:"password"`
	ServiceType  string `json:"service_type" db:"service_type"`
	Location     string `json:"location" db:"location"`
	Info         string `json:"info" db:"info"`
}

type UpdateUserInput struct {
	Name        *string `json:"name"`
	Username    *string `json:"username"`
	ServiceType *string `json:"service_type"`
	Location    *string `json:"location" db:"location"`
	Info        *string `json:"info"`
}

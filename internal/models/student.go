package models

import "time"

type Student struct {
	Id          int
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

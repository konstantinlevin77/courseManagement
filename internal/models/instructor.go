package models

import "time"

type Instructor struct {
	Id          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

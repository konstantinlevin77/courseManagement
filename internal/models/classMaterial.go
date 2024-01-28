package models

import "time"

type ClassMaterial struct {
	Id          int
	ClassId     int
	Type        string
	Description string
	Path        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

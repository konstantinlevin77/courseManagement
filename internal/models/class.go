package models

import "time"

type Class struct {
	Id          int
	CourseId    int
	Name        string
	Date        time.Time
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

package models

import "time"

type Class struct {
	Id          int       `json:"id"`
	CourseId    int       `json:"course_id"`
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

package models

import "time"

type InstructorAccess struct {
	Id           int
	InstructorId int
	CourseId     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

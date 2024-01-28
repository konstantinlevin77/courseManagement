package models

import "time"

type StudentAccess struct {
	Id        int
	StudentId int
	CourseId  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

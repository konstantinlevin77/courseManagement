package repository

import "github.com/konstantinlevin77/courseManagement/internal/models"

type DatabaseRepo interface {
	GetStudentById(studentId int) (models.Student, error)
	GetCourseById(courseId int) (models.Course, error)
}

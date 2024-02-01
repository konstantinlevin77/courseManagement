package repository

import "github.com/konstantinlevin77/courseManagement/internal/models"

type DatabaseRepo interface {
	GetStudentById(studentId int) (models.Student, error)

	GetCourseById(courseId int) (models.Course, error)
	GetCoursesByStudentId(studentId int) ([]models.Course, error)

	GetClassesByCourseId(courseId int) ([]models.Class, error)
}

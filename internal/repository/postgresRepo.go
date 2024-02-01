package repository

import (
	"context"
	"database/sql"
	"github.com/konstantinlevin77/courseManagement/internal/driver"
	"github.com/konstantinlevin77/courseManagement/internal/models"
	"time"
)

type PostgresRepo struct {
	DB *sql.DB
}

// GetCourseById returns the course whose id is passed as an argument.
func (pr *PostgresRepo) GetCourseById(courseId int) (models.Course, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT name,description,created_at,updated_at
              FROM course 
			  WHERE id=$1`

	row := pr.DB.QueryRowContext(ctx, query, courseId)

	var course models.Course
	course.Id = courseId

	err := row.Scan(&course.Name,
		&course.Description,
		&course.CreatedAt,
		&course.UpdatedAt,
	)

	return course, err

}

// GetStudentById returns the student whose id is passed as an argument.
func (pr *PostgresRepo) GetStudentById(studentId int) (models.Student, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT s.first_name, s.last_name, s.email, s.phone_number, s.password, s.created_at, s.updated_at
              FROM student s
              WHERE s.id=$1`

	row := pr.DB.QueryRowContext(ctx, query, studentId)
	var student models.Student
	student.Id = studentId

	err := row.Scan(&student.FirstName,
		&student.LastName,
		&student.Email,
		&student.PhoneNumber,
		&student.Password,
		&student.CreatedAt,
		&student.UpdatedAt,
	)

	return student, err

}

// GetCoursesByStudentId returns the courses that the student whose id is passed as an argument have access.
func (pr *PostgresRepo) GetCoursesByStudentId(studentId int) ([]models.Course, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT course_id FROM student_access WHERE student_id=$1`

	var courses []models.Course

	// Execute the query.
	rows, err := pr.DB.QueryContext(ctx, query, studentId)
	if err != nil {
		return courses, err
	}

	// Iterate through rows and add find corresponding courses by the next id.
	for rows.Next() {
		var nextId int
		err := rows.Scan(&nextId)
		if err != nil {
			continue
			// TODO: Error handling is missing.
		}

		c, err := pr.GetCourseById(nextId)
		if err != nil {
			continue
			// TODO: Error handling is missing.
		}

		courses = append(courses, c)
	}

	return courses, nil

}

// GetClassesByCourseId returns classes that belong to the course whose id is passed as an argument.
func (pr *PostgresRepo) GetClassesByCourseId(courseId int) ([]models.Class, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var classes []models.Class

	query := `SELECT id,name,date,description,created_at,updated_at FROM class WHERE course_id=$1`

	rows, err := pr.DB.QueryContext(ctx, query, courseId)
	if err != nil {
		return classes, err
	}

	for rows.Next() {

		var class models.Class
		class.CourseId = courseId
		err = rows.Scan(&class.Id, &class.Name, &class.Date, &class.Description, &class.CreatedAt, &class.UpdatedAt)
		if err != nil {
			continue
			// TODO: Error handling is missing.
		}

		classes = append(classes, class)

	}
	return classes, nil
}

func NewPostgresRepo(dsn string) (PostgresRepo, error) {

	db, err := driver.NewDB(dsn)
	if err != nil {
		return PostgresRepo{}, err
	}

	return PostgresRepo{DB: db}, nil

}

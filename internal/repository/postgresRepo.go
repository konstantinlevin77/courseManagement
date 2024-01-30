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

func NewPostgresRepo(dsn string) (PostgresRepo, error) {

	db, err := driver.NewDB(dsn)
	if err != nil {
		return PostgresRepo{}, err
	}

	return PostgresRepo{DB: db}, nil

}

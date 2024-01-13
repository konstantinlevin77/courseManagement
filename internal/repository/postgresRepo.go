package repository

import (
	"database/sql"
	"github.com/konstantinlevin77/courseManagement/internal/driver"
)

type PostgresRepo struct {
	DB *sql.DB
}

func (pr *PostgresRepo) Test() bool {
	return true
}

func NewPostgresRepo(dsn string) (PostgresRepo, error) {

	db, err := driver.NewDB(dsn)
	if err != nil {
		return PostgresRepo{}, err
	}

	return PostgresRepo{DB: db}, nil

}

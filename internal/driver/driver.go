package driver

import "database/sql"

// NewDB returns a new *sql.DB instance to operate the database.
func NewDB(dsn string) (*sql.DB, error) {

	// Postgres is used in this project, therefore the driver is pgx.
	driver, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = TestDB(driver)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

// TestDB pings the given database and returns an error if any.
func TestDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}

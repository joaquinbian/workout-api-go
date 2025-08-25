package store

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func setupDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable")

	if err != nil {
		t.Fatalf("Connect to DB: %v", err)
	}

	err = Migrate(db, "../../migrations")

	if err != nil {
		t.Fatalf("Running migrations: %v", err)
	}

	_, err = db.Exec("TRUNCATE workouts, workout_entries")

	if err != nil {
		t.Fatalf("Truncating tables: %v", err)
	}

	return db
}

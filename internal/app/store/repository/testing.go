package repository

import (
	"database/sql"
	"strings"
	"testing"
)

func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatalf("cannot open db: %s", err.Error())
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("error while ping db: %s", err.Error())
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec("TRUNCATE %s CASCADE", strings.Join(tables, ","))
		}

		db.Close()
	}
}
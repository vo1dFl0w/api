package repository_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("test_databaseurl")
	if databaseURL == "" {
		databaseURL = "host=localhost port=5433 user=admin password=admin dbname=wallets_db_test sslmode=disable"
	}

	os.Exit(m.Run())
}

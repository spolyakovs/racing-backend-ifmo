package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

// TODO: create tables and fill them before tests
// TODO: drop tables after tests

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=racing_ifmo_db_test user=racing_ifmo_user password=racing_ifmo123 sslmode=disable"
	}

	os.Exit(m.Run())
}

package db

import (
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
)

var testQueries *Queries

var mock sqlmock.Sqlmock

// openTestDB opens a new database connection that is not a pool
func openTestDB() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://root:myexamplepassword@localhost:5432/tweets?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

// TestUserMain runs the user test
func TestMain(m *testing.M) {
	useLiveDB := os.Getenv("USE_LIVE_DB")
	if useLiveDB != "" {
		// Setup live database connection
		db := openTestDB()
		defer db.Close()
		testQueries = New(db)
	} else {
		db, tmpMock, err := sqlmock.New()
		mock = tmpMock
		if err != nil {
			panic(err)
		}
		defer db.Close()
	}
	os.Exit(m.Run())
}

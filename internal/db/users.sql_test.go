// FILEPATH: /Users/chris/Documents/Go/1/internal/db/users_test.sql.go
// depends on users.sql.go, models.go

package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"example.com/utils"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username:  sql.NullString{String: utils.RandomUsername(), Valid: true},
		Role:      sql.NullString{String: utils.RandomRole(), Valid: true},
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	// create a new postgresl sql context
	testContext := context.Background()

	_, err := testQueries.CreateUser(testContext, arg)
	if err != nil {
		t.Fatal(err)
	}
}

// TestGetUser test the GetUser function
func TestGetUser(t *testing.T) {
	arg := CreateUserParams{
		Username:  sql.NullString{String: utils.RandomUsername(), Valid: true},
		Role:      sql.NullString{String: utils.RandomRole(), Valid: true},
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	// create a new postgresl sql context
	testContext := context.Background()

	// create a new user
	user, err := testQueries.CreateUser(testContext, arg)
	if err != nil {
		t.Fatal(err)
	}

	// get the user
	user, err = testQueries.GetUser(testContext, user.ID)
	if err != nil {
		t.Fatal(err)
	}

	if user.Username.String != arg.Username.String {
		t.Errorf("got %v, expected %v", user.Username, arg.Username.String)
	}
}

// TestDeleteUser test the DeleteUser function
func TestDeleteUser(t *testing.T) {
	arg := CreateUserParams{
		Username:  sql.NullString{String: utils.RandomUsername(), Valid: true},
		Role:      sql.NullString{String: utils.RandomRole(), Valid: true},
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	// create a new postgresl sql context
	testContext := context.Background()

	// create a new user
	user, err := testQueries.CreateUser(testContext, arg)
	if err != nil {
		t.Fatal(err)
	}

	// delete the user
	err = testQueries.DeleteUser(testContext, user.ID)
	if err != nil {
		t.Fatal(err)
	}
}

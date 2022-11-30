package users_test

import (
	"assignment-activity-reporter/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabase_AddUser(t *testing.T) {
	t.Run("Should return no error when a new user is inputted to the database", func(t *testing.T) {
		alice := users.NewUser("Alice")
		database := users.NewDB()

		err := database.AddUser(alice)
		assert.Nil(t, err)
	})

	t.Run("Should return UserAlreadyExists error when a user is inputted to the database twice", func(t *testing.T) {
		alice := users.NewUser("Alice")
		database := users.NewDB()

		database.AddUser(alice)
		err := database.AddUser(alice)
		assert.Exactly(t, err, &users.UserAlreadyExistError{})
	})
}

func TestDatabase_GetUser(t *testing.T) {
	t.Run("Should return correct user when given a key that already exists", func(t *testing.T) {
		alice := users.NewUser("Alice")
		database := users.NewDB()

		database.AddUser(alice)
		result, err := database.GetUser("Alice")
		assert.Nil(t, err)
		assert.Exactly(t, result, alice)
	})
}

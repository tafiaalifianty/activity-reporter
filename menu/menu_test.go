package menu

import (
	"assignment-activity-reporter/users"
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMenu(t *testing.T) {
	t.Run("Menu text should be equal to formatted text below", func(t *testing.T) {
		menu := *GetMenu()
		expected := "Activity Reporter\n" +
			"1. Setup\n" +
			"2. Action\n" +
			"3. Display\n" +
			"4. Trending\n" +
			"5. Exit"
		assert.Equal(t, menu, expected)
	})
}

func TestExitMessage(t *testing.T) {
	t.Run("Exit message should be equal to the expected output", func(t *testing.T) {
		message := *ExitMessage()
		expected := "Good Bye!"

		assert.Equal(t, message, expected)
	})
}

func TestPromptInput(t *testing.T) {
	t.Run("Scanner should be able to read input text correctly", func(t *testing.T) {
		var r io.Reader = strings.NewReader("Hello")

		scanner := bufio.NewScanner(r)
		input := CueInput(scanner, "Enter user actions: ")
		fmt.Println(input)
		assert.Equal(t, "Hello", input)
	})
}

func TestSetup(t *testing.T) {
	t.Run("Should add 2 users to database when correct syntax is given", func(t *testing.T) {
		var r io.Reader = strings.NewReader("Alice follows Bob")
		scanner := bufio.NewScanner(r)
		database := users.NewDB()
		Setup(database, scanner)
		assert.Equal(t, len(database.User), 2)
	})
	t.Run("Database should remain empty when incorrect syntax is given", func(t *testing.T) {
		var r io.Reader = strings.NewReader("Alice says hi to Bob")
		scanner := bufio.NewScanner(r)
		database := users.NewDB()
		Setup(database, scanner)
		assert.Equal(t, len(database.User), 0)
	})
}

func TestAction(t *testing.T) {
	t.Run("Bob should be able to upload photo after registering himself", func(t *testing.T) {
		var r io.Reader = strings.NewReader("Alice follows Bob")
		scanner := bufio.NewScanner(r)
		database := users.NewDB()
		Setup(database, scanner)

		var r2 io.Reader = strings.NewReader("Bob uploaded photo")
		scanner2 := bufio.NewScanner(r2)
		Action(database, scanner2)
		assert.NotNil(t, database.User["Bob"].Photos)
		assert.Equal(t, database.User["Bob"].Activities[0], "You uploaded photo")
	})

	t.Run("Alice should be able to like Bob's photo given 2 conditions:", func(t *testing.T) {
		//"1. Alice follows Bob", "2. Bob has uploaded a photo before"
		var r io.Reader = strings.NewReader("Alice follows Bob")
		scanner := bufio.NewScanner(r)
		database := users.NewDB()
		Setup(database, scanner)

		var r2 io.Reader = strings.NewReader("Bob uploaded photo")
		scanner2 := bufio.NewScanner(r2)
		Action(database, scanner2)

		var r3 io.Reader = strings.NewReader("Alice likes Bob photo")
		scanner3 := bufio.NewScanner(r3)
		Action(database, scanner3)

		assert.Equal(t, database.User["Alice"].Activities[1], "You liked Bob's photo")
		assert.Equal(t, database.User["Bob"].Activities[1], "Alice liked your photo")
	})
}

func TestDisplay(t *testing.T) {
	t.Run("Should be able to display correct activities for multiple people", func(t *testing.T) {
		var r io.Reader = strings.NewReader("Alice follows Bob")
		scanner := bufio.NewScanner(r)
		database := users.NewDB()
		Setup(database, scanner)

		var r2 io.Reader = strings.NewReader("Bob uploaded photo")
		scanner2 := bufio.NewScanner(r2)
		Action(database, scanner2)

		var r3 io.Reader = strings.NewReader("Alice likes Bob photo")
		scanner3 := bufio.NewScanner(r3)
		Action(database, scanner3)

		var r4 io.Reader = strings.NewReader("Alice")
		scanner4 := bufio.NewScanner(r4)
		Display(database, scanner4)

		assert.Equal(t, database.User["Alice"].Activities[0], "Bob uploaded photo")
		assert.Equal(t, database.User["Alice"].Activities[1], "You liked Bob's photo")
		assert.Equal(t, database.User["Bob"].Activities[0], "You uploaded photo")
		assert.Equal(t, database.User["Bob"].Activities[1], "Alice liked your photo")
	})
}

func TestTrending(t *testing.T) {
	t.Run("when no user has upload photo, should return NoUserUploadedError", func(t *testing.T) {
		// Given
		system := users.NewDB()

		// When
		err := Trending(system)

		// Then
		if assert.Error(t, err) {
			assert.ErrorContains(t, err, "No user has uploaded photo")
		}
	})

	t.Run("when user(s) has upload photo, should not return err", func(t *testing.T) {
		// Given
		system := users.NewDB()
		username1 := "user1"
		user1 := users.NewUser(username1)
		system.AddUser(user1)
		user1.Upload(&users.Photo{})

		// When
		err := Trending(system)

		// Then
		assert.NoError(t, err)
	})
}

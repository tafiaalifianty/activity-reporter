package inputchecker_test

import (
	inputchecker "assignment-activity-reporter/inputChecker"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFollowsChecker_CheckInput(t *testing.T) {
	t.Run("Should return InvalidLengthError if input command is not exactly 3 words", func(t *testing.T) {
		command := "Alice follows"
		inputChecker := new(inputchecker.CheckFollows)

		err := inputChecker.CheckInput(&command)
		assert.Exactly(t, err, &inputchecker.InvalidLengthError{Length: 3})
	})
	t.Run("Should return NoKeywordError if input command does not contains: follows", func(t *testing.T) {
		command := "Alice following Bob"
		inputChecker := new(inputchecker.CheckFollows)

		err := inputChecker.CheckInput(&command)
		assert.Exactly(t, err, &inputchecker.NoKeywordError{Keyword: "follows"})
	})
	t.Run("Should return no error if input contains 'follows' keyword and has 3 words", func(t *testing.T) {
		command := "Alice follows Bob"
		inputChecker := new(inputchecker.CheckFollows)

		err := inputChecker.CheckInput(&command)
		assert.Nil(t, err)
	})
}

func TestUploadedChecker_CheckInput(t *testing.T) {
	t.Run("Should return InvalidLengthError if input command is not exactly 3 words", func(t *testing.T) {
		command := "Alice uploaded"
		inputChecker := new(inputchecker.CheckUpload)

		err := inputChecker.CheckInput(&command)
		assert.Exactly(t, err, &inputchecker.InvalidLengthError{Length: 3})
	})
	t.Run("Should return NoKeywordError if the second word is not 'uploaded'", func(t *testing.T) {
		command := "Alice uploads photo"
		inputChecker := new(inputchecker.CheckUpload)

		err := inputChecker.CheckInput(&command)
		assert.Exactly(t, err, &inputchecker.NoKeywordError{Keyword: "uploaded"})
	})
	t.Run("Should return NoKeywordError if the third word is not 'photo'", func(t *testing.T) {
		command := "Alice uploaded whatever"
		inputChecker := new(inputchecker.CheckUpload)

		err := inputChecker.CheckInput(&command)
		assert.Exactly(t, err, &inputchecker.NoKeywordError{Keyword: "photo"})
	})
	t.Run("Should return no error if the format is as follows: (name) uploaded photo", func(t *testing.T) {
		command := "Alice uploaded photo"
		inputChecker := new(inputchecker.CheckUpload)

		err := inputChecker.CheckInput(&command)
		assert.Nil(t, err)
	})
}

func TestLikesChecker_CheckInput(t *testing.T) {
	t.Run("Should return InvalidLengthError if input command is not exactly 4 words", func(t *testing.T) {
		command := "Alice likes Bob"
		inputChecker := new(inputchecker.CheckLikes)

		err := inputChecker.CheckInput(&command)
		assert.Exactly(t, err, &inputchecker.InvalidLengthError{Length: 4})
	})
	t.Run("Should return NoKeywordError if the second word is not 'likes'", func(t *testing.T) {
		command := "Alice loves Bob photo"
		inputChecker := new(inputchecker.CheckLikes)

		err := inputChecker.CheckInput(&command)
		assert.Exactly(t, err, &inputchecker.NoKeywordError{Keyword: "likes"})
	})
	t.Run("Should return NoKeywordError if the fourth word is not 'photo'", func(t *testing.T) {
		command := "Alice likes Bob pants"
		inputChecker := new(inputchecker.CheckLikes)

		err := inputChecker.CheckInput(&command)
		assert.Exactly(t, err, &inputchecker.NoKeywordError{Keyword: "photo"})
	})
	t.Run("Should return no error if the format is as follows: (name) likes (name) photo", func(t *testing.T) {
		command := "Alice likes Bob photo"
		inputChecker := new(inputchecker.CheckLikes)

		err := inputChecker.CheckInput(&command)
		assert.Nil(t, err)
	})
}

func TestDisplayChecker_CheckInput(t *testing.T) {
	t.Run("Should return InvalidLengthError if input command is not exactly 1 word", func(t *testing.T) {
		command := "Alice in wonderland"
		inputChecker := new(inputchecker.CheckDisplay)

		err := inputChecker.CheckInput(&command)
		assert.Exactly(t, err, &inputchecker.InvalidLengthError{Length: 1})
	})
	t.Run("Should return no error if the format is as follows: (name)", func(t *testing.T) {
		command := "Alice"
		inputChecker := new(inputchecker.CheckDisplay)

		err := inputChecker.CheckInput(&command)
		assert.Nil(t, err)
	})
}

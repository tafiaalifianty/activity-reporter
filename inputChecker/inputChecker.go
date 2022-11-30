package inputchecker

import (
	"fmt"
	"strings"
)

type CheckFollows struct {
}

type CheckUpload struct {
}

type CheckLikes struct {
}

type CheckDisplay struct {
}

type InvalidLengthError struct {
	Length int
}

type NoKeywordError struct {
	Keyword string
}

type NoUserUploadedError struct{}

func (e NoUserUploadedError) Error() string {
	return "No user has uploaded photo"
}

func (il *InvalidLengthError) Error() string {
	message := fmt.Sprintf("Input should contain %v word", il.Length)

	return message
}

func (nk *NoKeywordError) Error() string {
	message := fmt.Sprintf("Input doesn't contain %v keyword", nk.Keyword)
	return message
}

func (cf *CheckFollows) CheckInput(s *string) error {
	reqLength := 3
	keyword := map[string]int{
		"follows": 2,
	}
	return ErrorHandling(reqLength, &keyword, s)
}

func (cu *CheckUpload) CheckInput(s *string) error {
	reqLength := 3
	keyword := map[string]int{
		"uploaded": 2,
		"photo":    3,
	}
	return ErrorHandling(reqLength, &keyword, s)
}

func (cl *CheckLikes) CheckInput(s *string) error {
	reqLength := 4
	keyword := map[string]int{
		"likes": 2,
		"photo": 4,
	}
	return ErrorHandling(reqLength, &keyword, s)
}

func (cd *CheckDisplay) CheckInput(s *string) error {
	reqLength := 1
	keyword := map[string]int{}

	return ErrorHandling(reqLength, &keyword, s)
}

func ErrorHandling(reqLength int, keyword *map[string]int, input *string) error {
	words := strings.Split(*input, " ")
	if len(words) != reqLength {
		return &InvalidLengthError{Length: reqLength}
	}
	for keyword, wordPos := range *keyword {
		if words[wordPos-1] != keyword {
			return &NoKeywordError{Keyword: keyword}
		}
	}
	return nil
}

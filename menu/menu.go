package menu

import (
	inputchecker "assignment-activity-reporter/inputChecker"
	"assignment-activity-reporter/users"
	"bufio"
	"fmt"
	"strings"
)

func GetMenu() *string {
	menu := "Activity Reporter\n" +
		"1. Setup\n" +
		"2. Action\n" +
		"3. Display\n" +
		"4. Trending\n" +
		"5. Exit"

	return &menu
}

func CueInput(scanner *bufio.Scanner, text string) string {
	fmt.Print(text)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func Setup(db *users.Database, scanner *bufio.Scanner) {
	input := CueInput(scanner, "Setup social graph: ")
	inputChecker := new(inputchecker.CheckFollows)
	err := inputChecker.CheckInput(&input)
	if err == nil {
		words := strings.Split(input, " ")
		followers := words[0]
		followed := words[2]

		if !db.IsExist(followers) {
			newUser := users.NewUser(followers)
			db.AddUser(newUser)
		}

		if !db.IsExist(followed) {
			newUser2 := users.NewUser(followed)
			db.AddUser(newUser2)
		}

		dbFollowers, _ := db.GetUser(followers)
		dbFollowed, _ := db.GetUser(followed)
		dbFollowers.Follow(dbFollowed)
	} else {
		fmt.Println(err)
	}
	fmt.Println("---------------------------------------")
}

func Action(db *users.Database, scanner *bufio.Scanner) {
	input := CueInput(scanner, "Enter User Actions: ")
	uploadChecker := new(inputchecker.CheckUpload)
	likedChecker := new(inputchecker.CheckLikes)
	errUpload := uploadChecker.CheckInput(&input)
	errLike := likedChecker.CheckInput(&input)

	if errUpload != nil && errLike != nil {
		fmt.Println(errUpload)
		fmt.Println(errLike)
	} else {
		words := strings.Split(input, " ")
		if errUpload == nil {
			if user, err := db.GetUser(words[0]); err == nil {
				user.Upload(new(users.Photo))
			} else {
				fmt.Println(users.NoUserExistError{})
			}
		}
		if errLike == nil {
			likers, err := db.GetUser(words[0])
			liked, err2 := db.GetUser(words[2])

			if err == nil && err2 == nil && liked.Photos != nil {
				likers.LikePhoto(liked)
			} else {
				fmt.Println(users.NoUserExistError{})
			}
		}
	}
	fmt.Println("---------------------------------------")
}

func Display(db *users.Database, scanner *bufio.Scanner) {
	input := CueInput(scanner, "Display Activity for: ")
	inputChecker := new(inputchecker.CheckDisplay)
	err := inputChecker.CheckInput(&input)

	if err == nil {
		if user, err := db.GetUser(input); err == nil {
			for _, activity := range user.Activities {
				fmt.Println(activity)
			}
		} else {
			fmt.Println(users.NoUserExistError{})
		}
	} else {
		fmt.Println(err)
	}
	fmt.Println("---------------------------------------")
}

func Trending(sys *users.Database) error {
	users := sys.GetTrendingUser()

	if len(users) == 0 {
		return inputchecker.NoUserUploadedError{}
	}
	fmt.Printf("\nTrending Photos: ")

	for i, user := range users {
		fmt.Printf("\n%d. %s photo got %d likes\n", i+1, user.GetName(), user.GetLikersCount())
	}
	return nil

}

func ExitMessage() *string {
	message := "Good Bye!"

	return &message
}

type InvalidMenuError struct {
}

func (i *InvalidMenuError) Error() string {
	return "Invalid Menu"
}

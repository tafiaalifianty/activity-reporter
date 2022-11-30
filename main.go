package main

import (
	"assignment-activity-reporter/menu"
	"assignment-activity-reporter/users"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	menuDisplay := *menu.GetMenu()
	exit := false
	db := users.NewDB()
	fmt.Println(menuDisplay)

	for {
		input := menu.CueInput(scanner, "Enter Menu: ")

		switch input {
		case "1":
			menu.Setup(db, scanner)
		case "2":
			menu.Action(db, scanner)
		case "3":
			menu.Display(db, scanner)
		case "4":
			menu.Trending(db)
		case "5":
			fmt.Println(*menu.ExitMessage())
			exit = true
		default:
			err := menu.InvalidMenuError{}
			fmt.Println(err.Error())
		}

		if exit {
			break
		}
	}
}

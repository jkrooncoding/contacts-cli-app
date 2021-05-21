package main

import (
	"fmt"
	"personal/contactscli/models"
	"strings"
)

var asciiTitle string = `
 _____          _____             _             _       
/ ____|        / ____|           | |           | |      
| |  __  ___   | |     ___  _ __ | |_ __ _  ___| |_ ___ 
| | |_ |/ _ \  | |    / _ \| '_ \| __/ _' |/ __| __/ __|
| |__| | (_) | | |___| (_) | | | | || (_| | (__| |_\__ \
\_____ |\___/   \_____\___/|_| |_|\__\__,_|\___|\__|___/`

var contacts []models.Contact
var quit bool = false

func main() {
	readContactsFromFile()
	for !quit {
		clearScreen()
		printHeader("")
		printMenu()
		userChoice, err := getValidationInput(InputTypeMenuItem)
		if err != nil {
			errorString := err.Error()
			fmt.Println()
			fmt.Println(strings.Title(errorString))
			fmt.Println("Press enter to try again.")
			getInput()
		} else {
			option := menuItems[userChoice]
			runOption(option)
		}
	}
}

// runOption takes in a option and, clears the screen, prints the header with
// the option's description and then runs its function
func runOption(o models.Option) {
	clearScreen()
	printHeader(o.Description)
	o.Function()
}

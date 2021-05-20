package main

import (
	"log"
)

var contacts []contact
var quit bool = false

func main() {
	readContactsFromFile()
	for !quit {
		clearScreen()
		printHeader("")
		menu()
		userChoice, err := getValidationInput(InputTypeMenuItem)
		if err != nil {
			log.Fatal(err)
		}
		option := menuItems[userChoice]
		runOption(option)
	}
}

// runOption takes in a option and, clears the screen, prints the header with
// the option's description and then runs its function
func runOption(o option) {
	clearScreen()
	printHeader(o.description)
	o.function()
}

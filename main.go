package main

import (
	"log"
)

var contacts []contact
var quit bool = false

func main() {
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

func runOption(o option) {
	clearScreen()
	printHeader(o.description)
	o.function()
}

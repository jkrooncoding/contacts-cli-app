package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/inancgumus/screen"
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

func printHeader(description string) {
	fmt.Println("////////////////////////////////")
	fmt.Println("/////// GO CONTACTS V0.1 ///////")
	fmt.Println("////////////////////////////////")
	fmt.Println("\t", strings.ToUpper(description))
	fmt.Println("")
}

func clearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}

func quitProgram() {
	fmt.Println("Quitting GO Contacts...")
	quit = true
	os.Exit(0)
}

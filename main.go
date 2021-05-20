package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/inancgumus/screen"
)

var contacts []contact

func main() {
	printHeader("")
	menu()
	userChoice, err := getValidationInput(InputTypeMenuItem)
	if err != nil {
		log.Fatal(err)
	}
	option := menuItems[userChoice]
	runOption(option)
}

func runOption(o option) {
	screen.Clear()
	screen.MoveTopLeft()
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

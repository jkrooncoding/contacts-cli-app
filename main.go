package main

import (
	"fmt"
	"log"
)

// var contacts []contact

func main() {
	fmt.Println("Welcome to GoContacts v0.1")
	menu()
	userChoice, err := getInput()
	if err != nil {
		log.Fatal(err)
	}
	option := menuItems[userChoice]
	runOption(option)
}

func runOption(o option) {
	o.function()
}

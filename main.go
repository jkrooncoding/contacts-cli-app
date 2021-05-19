package main

import (
	"log"
)

var contacts []contact

func main() {
	menu()
	userChoice, err := getInput()
	if err != nil {
		log.Fatal(err)
	}
	option := menuItems[userChoice]
	runOption(option)
}

func runOption(o option) {
	switch 
}

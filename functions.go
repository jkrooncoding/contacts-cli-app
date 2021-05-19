package main

import (
	"fmt"
)

func AddContact() {
	fmt.Print("Please enter your first name: ")
	firstName := getInput()

	fmt.Print("Please enter  your second name: ")
	lastName := getInput()

	fmt.Print("Please enter email address: ")
	email := getInput()

	fmt.Print("Please enter phone number: ")
	phoneNumber := getInput()

	fmt.Println(firstName, lastName)
}

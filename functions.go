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

	newContact := contact{
		firstName:    firstName,
		lastName:     lastName,
		emailAddress: email,
		phoneNumber:  phoneNumber,
	}

	contacts = append(contacts, newContact)

	fmt.Println(contacts)
}

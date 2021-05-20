package main

import (
	"fmt"
	"strings"
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
		FirstName:    strings.ToLower(firstName),
		LastName:     strings.ToLower(lastName),
		EmailAddress: email,
		PhoneNumber:  phoneNumber,
	}

	contacts = append(contacts, newContact)

	saveToFile()
}

func ViewContact() {
	fmt.Println("Enter a full name to search for a specific contact")
	fmt.Println("or enter a last name to search for matches")
	fmt.Print("Please enter searchterm: ")
	contactName := getInput()
	if len(strings.Split(contactName, " ")) == 2 {
		searchSpecificContact(contactName)
	} else {
		searchByLastName(contactName)
	}
}

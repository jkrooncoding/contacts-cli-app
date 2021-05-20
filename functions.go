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
	fmt.Print("Please enter contact's name: ")
	contactName := getInput()
	contactFirstName := strings.Split(contactName, " ")[0]
	contactLastName := strings.Split(contactName, " ")[1]

	for _, c := range contacts {
		if c.FirstName == contactFirstName && c.LastName == contactLastName {
			printContact(c)
			if getInput() == "" {
				return
			}
		} else {
			continue
		}
	}

	fmt.Println("No contact found. Please check your spelling.")
	if getInput() == "" {
		return
	}
}

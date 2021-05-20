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
		firstName:    strings.ToLower(firstName),
		lastName:     strings.ToLower(lastName),
		emailAddress: email,
		phoneNumber:  phoneNumber,
	}

	contacts = append(contacts, newContact)

	fmt.Println(contacts)
}

func ViewContact() {
	fmt.Print("Please enter contact's name: ")
	contactName := getInput()
	contactFirstName := strings.Split(contactName, " ")[0]
	contactLastName := strings.Split(contactName, " ")[1]

	for _, c := range contacts {
		if c.firstName == contactFirstName && c.lastName == contactLastName {
			printContact(c)
			if getInput() == "" {
				return
			}
		} else {
			fmt.Println("No contact found. Please check your spelling.")
			if getInput() == "" {
				return
			}
		}
	}
}

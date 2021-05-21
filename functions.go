package main

import (
	"fmt"
	"personal/contactscli/models"
	"strings"
)

func AddContact() {
	fmt.Print("Please enter first name: ")
	firstName := getInput()

	fmt.Print("Please enter last name: ")
	lastName := getInput()

	fmt.Print("Please enter email address: ")
	email := getInput()

	fmt.Print("Please enter phone number: ")
	phoneNumber := getInput()

	newContact := models.Contact{
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
	fmt.Print("Please enter search term: ")
	contactName := getInput()
	if len(strings.Split(contactName, " ")) == 2 {
		searchSpecificContact(contactName, "view")
	} else {
		searchByLastName(contactName)
	}
}

func EditContact() {
	//fmt.Println("Enter the full name of the contact you wish to edit")
	//fmt.Print("Enter contact's name: ")
	//contactName := getInput()
}

func DeleteContact() {
	fmt.Print("Enter the full name of the contact you wish to delete: ")
	contactName := getInput()
	if len(strings.Split(contactName, " ")) != 2 {
		fmt.Println("We only got one name. Deletion of contact requires full name.")
		getInput()
		return
	}
	searchSpecificContact(contactName, "delete")
}
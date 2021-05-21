package main

import (
	"encoding/json"
	"fmt"
	"os"
	"personal/contactscli/models"
	"strings"
	"time"

	"github.com/inancgumus/screen"
)

// printHeader prints out the header for our Go Contacts program
// including the description passed into it
func printHeader(description string) {
	fmt.Println(asciiTitle)
	fmt.Println("\t", strings.ToUpper(description))
	fmt.Println("")
}

// clearScreen uses the 'screen' package to get a blank screen
// and move the cursor to the top left of the terminal
func clearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}

// quitProgram exits the main for loop and also exits
// the program so that the rest of main doesn't
// continue running
func quitProgram() {
	fmt.Println("Quitting GO Contacts...")
	time.Sleep(2 * time.Second)
	quit = true
	os.Exit(0)
}

// printContact prints a contact's information
func printContact(c models.Contact) {
	fullName := strings.Join([]string{c.FirstName, c.LastName}, " ")
	fmt.Println("Name:\t", strings.Title(fullName))
	fmt.Println("Email:\t", c.EmailAddress)
	fmt.Println("Phone:\t", c.PhoneNumber)
}

// saveToFile saves all contacts to a file called contacts.json
func saveToFile() {
	file, _ := json.MarshalIndent(contacts, "", " ")
	err := os.WriteFile("contacts.json", file, 0666)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func readContactsFromFile() {
	byteContacts, err := os.ReadFile("contacts.json")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = json.Unmarshal(byteContacts, &contacts)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func searchMatches(search string) []models.Contact {
	var matches []models.Contact

	for _, c := range contacts {
		if c.LastName == search {
			matches = append(matches, c)
		} else {
			continue
		}
	}

	return matches
}

func searchByLastName(name string) {
	matches := searchMatches(name)

	clearScreen()
	printHeader("")

	if len(matches) == 0 {
		fmt.Println("No contacts found. Please check your spelling.")
		getInput()
		return
	}

	fmt.Println("These contacts match your search:")
	fmt.Println("")
	for _, c := range matches {
		printContact(c)
		fmt.Println()
	}
	getInput()
}

func searchSpecificContact(name, option string) {
	var match models.Contact
	contactFirstName := strings.Split(name, " ")[0]
	contactLastName := strings.Split(name, " ")[1]

	clearScreen()
	printHeader("")

	for _, c := range contacts {
		if c.FirstName == contactFirstName && c.LastName == contactLastName {
			match = c

			switch option {
			case "view":
				handleView(match)
			case "edit":
				handleEdit(match)
			case "delete":
				handleDelete(match)
			}
		} else {
			continue
		}
	}

	if match == (models.Contact{}) {
		fmt.Println("No contacts found. Please check your spelling.")
		getInput()
		return
	}
}

func handleView(c models.Contact) {
	printContact(c)
	getInput()
}

func handleEdit(c models.Contact) {
	printContact(c)

	fmt.Println("\nPlease enter new information. Leave blank to use existing value.")
	fmt.Println()
	fmt.Print("Please enter first name: ")
	firstName := getInput()
	if firstName == "" {
		firstName = c.FirstName
	}

	fmt.Print("Please enter last name: ")
	lastName := getInput()
	if lastName == "" {
		lastName = c.LastName
	}

	fmt.Print("Please enter email address: ")
	email := getInput()
	if email == "" {
		email = c.EmailAddress
	}

	fmt.Print("Please enter phone number: ")
	phoneNumber := getInput()
	if phoneNumber == "" {
		phoneNumber = c.PhoneNumber
	}

	newContact := models.Contact{
		FirstName:    strings.ToLower(firstName),
		LastName:     strings.ToLower(lastName),
		EmailAddress: email,
		PhoneNumber:  phoneNumber,
	}

	editContact(c, newContact)
}

func handleDelete(c models.Contact) {
	printContact(c)
	fmt.Println("\nAre you sure you wish to delete this contact? (Y)es/(N)o")
	userChoice := getInput()
	if strings.ToLower(string(userChoice[0])) == "y" {
		removeFromContacts(c)
	}
	fmt.Println("Contact successfully deleted. Press enter to go back to the main menu.")
	getInput()
}

func editContact(old, new models.Contact) {
	for i, v := range contacts {
		if v == old {
			contacts[i] = new
		}
	}
	saveToFile()
}

func removeFromContacts(c models.Contact) {
	for i, v := range contacts {
		if v == c {
			contacts = append(contacts[0:i], contacts[i+1:]...)
		}
	}

	saveToFile()
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/inancgumus/screen"
)

// printHeader prints out the header for our Go Contacts program
// including the description passed into it
func printHeader(description string) {
	fmt.Println(ascititle)
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
func printContact(c contact) {
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

func searchMatches(search string) []contact {
	var matches []contact

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

func searchSpecificContact(name string) {
	var match contact
	contactFirstName := strings.Split(name, " ")[0]
	contactLastName := strings.Split(name, " ")[1]

	clearScreen()
	printHeader("")

	for _, c := range contacts {
		if c.FirstName == contactFirstName && c.LastName == contactLastName {
			match = c
			printContact(match)
			getInput()
			return
		} else {
			continue
		}
	}

	if match == (contact{}) {
		fmt.Println("No contacts found. Please check your spelling.")
		getInput()
		return
	}
}

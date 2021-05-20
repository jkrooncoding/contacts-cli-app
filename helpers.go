package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/inancgumus/screen"
)

// printHeader prints out the header for our Go Contacts program
// including the description passed into it
func printHeader(description string) {
	fmt.Println("////////////////////////////////")
	fmt.Println("/////// GO CONTACTS V0.1 ///////")
	fmt.Println("////////////////////////////////")
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

func saveToFile() {
	file, _ := json.MarshalIndent(contacts, "", " ")
	err := os.WriteFile("contacts.json", file, 0666)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

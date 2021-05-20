package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/inancgumus/screen"
)

func printHeader(description string) {
	fmt.Println("////////////////////////////////")
	fmt.Println("/////// GO CONTACTS V0.1 ///////")
	fmt.Println("////////////////////////////////")
	fmt.Println("\t", strings.ToUpper(description))
	fmt.Println("")
}

func clearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}

func quitProgram() {
	fmt.Println("Quitting GO Contacts...")
	quit = true
	os.Exit(0)
}

func printContact(c contact) {
	fullName := strings.Join([]string{c.firstName, c.lastName}, " ")
	fmt.Println("Name:\t", strings.Title(fullName))
	fmt.Println("Email:\t", c.emailAddress)
	fmt.Println("Phone:\t", c.phoneNumber)
}

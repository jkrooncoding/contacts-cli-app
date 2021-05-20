package main

import "fmt"

type option struct {
	description string
	function    func()
}

var add option = option{
	description: "Add a contact",
	function:    AddContact,
}

var view option = option{
	description: "View a contact's details",
	function:    ViewContact,
}

var placeHolder option = option{
	description: "Placeholder",
	function: func() {
		fmt.Println("test")
	},
}

package main

import "fmt"

type option struct {
	description string
	function    func()
}

var add = option{
	description: "Add a contact",
	function:    AddContact,
}

var view = option{
	description: "View a contact",
	function:    ViewContact,
}

var placeHolder option = option{
	description: "Placeholder",
	function: func() {
		fmt.Println("test")
	},
}

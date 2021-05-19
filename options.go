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

var placeHolder option = option{
	description: "Placeholder option",
	function: func() {
		fmt.Println("test")
	},
}

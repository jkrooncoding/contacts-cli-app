package main

import "fmt"

type option struct {
	description string
	function    func()
}

var addContact option = option{
	description: "Add a contact",
	function: func() {
		fmt.Println("Adding fake contact")
	},
}

var placeHolder option = option{
	description: "Placeholder option",
	function: func() {
		fmt.Println("test")
	},
}

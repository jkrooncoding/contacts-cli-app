package main

import (
	"fmt"
)

var add = Option{
	Description: "Add a contact",
	Function:    AddContact,
}

var view = Option{
	Description: "View a contact",
	Function:    ViewContact,
}

var edit = Option{
	Description: "Edit a contact",
	Function:    EditContact,
}

var delete = Option{
	Description: "Delete a contact",
	Function:    DeleteContact,
}

var placeHolder = Option{
	Description: "Placeholder",
	Function: func() {
		fmt.Println("test")
	},
}

package main

import (
	"fmt"
	"github.com/jkrooncoding/contacts-cli-app/models"
)

var add = models.Option{
	Description: "Add a contact",
	Function:    AddContact,
}

var view = models.Option{
	Description: "View a contact",
	Function:    ViewContact,
}

var edit = models.Option{
	Description: "Edit a contact",
	Function:    EditContact,
}

var delete = models.Option{
	Description: "Delete a contact",
	Function:    DeleteContact,
}

var placeHolder = models.Option{
	Description: "Placeholder",
	Function: func() {
		fmt.Println("test")
	},
}

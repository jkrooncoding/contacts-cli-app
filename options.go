package main

import (
	"fmt"
	"personal/contactscli/models"
)

var add = models.Option{
	Description: "Add a contact",
	Function:    AddContact,
}

var view = models.Option{
	Description: "View a contact",
	Function:    ViewContact,
}

var placeHolder = models.Option{
	Description: "Placeholder",
	Function: func() {
		fmt.Println("test")
	},
}

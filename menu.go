package main

import (
	"fmt"
	"personal/contactscli/models"
	"sort"
)

var menuItems map[string]models.Option = map[string]models.Option{
	"1": add,
	"2": placeHolder,
	"3": view,
	"4": placeHolder,
}

// menu prints out the program's main menu and all the possible
// options a user can enter
func printMenu() {
	sortedMenuItems := sortMenuItems(menuItems)
	for _, k := range sortedMenuItems {
		fmt.Printf("%s. %s\n", k, menuItems[k].Description)
	}
	fmt.Println("\nOr enter 'q' to exit the program.")

	fmt.Print("Your option: ")
}

// sortMenuItems adds the menuItems map to a slice of strings
// then sorts the slice and returns it
func sortMenuItems(menuItems map[string]models.Option) []string {
	keys := make([]string, 0)
	for k := range menuItems {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

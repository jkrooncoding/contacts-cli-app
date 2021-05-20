package main

import (
	"fmt"
	"sort"
)

var menuItems map[string]option = map[string]option{
	"1": add,
	"2": placeHolder,
	"3": placeHolder,
	"4": placeHolder,
	"5": placeHolder,
	"6": placeHolder,
}

func menu() {
	sortedMenuItems := sortMenuItems(menuItems)
	for _, k := range sortedMenuItems {
		fmt.Printf("%s. %s\n", k, menuItems[k].description)
	}
	fmt.Println("Or enter 'q' to exit the program.")

	fmt.Print("Your option: ")
}

func sortMenuItems(menuItems map[string]option) []string {
	keys := make([]string, 0)
	for k := range menuItems {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

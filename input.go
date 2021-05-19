package main

import (
	"bufio"
	"errors"
	"os"
)

func getInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	_, err := validateInput(scanner.Text())
	if err != nil {
		return "", err
	}

	return scanner.Text(), nil
}

func validateInput(input string) (string, error) {
	if menuItems[input].description != "" {
		return menuItems[input].description, nil
	} else {
		return "", errors.New("input is invalid")
	}
}

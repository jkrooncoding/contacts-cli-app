package main

import (
	"bufio"
	"errors"
	"os"
)

func getInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err := validateMenuInput(scanner.Text())
	if err != nil {
		return "", err
	}

	return scanner.Text(), nil
}

func validateMenuInput(input string) error {
	if menuItems[input].description != "" {
		return nil
	} else {
		return errors.New("input is invalid")
	}
}

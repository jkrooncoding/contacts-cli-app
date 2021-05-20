package main

import (
	"bufio"
	"errors"
	"os"
)

func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func getValidationInput(inputType InputType) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err := validateInput(scanner.Text(), inputType)
	if err != nil {
		return "", err
	}

	return scanner.Text(), nil
}

func validateInput(input string, inputType InputType) error {
	switch inputType {
	case InputTypeMenuItem:
		if input == "q" {
			quitProgram()
			return nil
		} else if menuItems[input].description != "" {
			return nil
		} else {
			return errors.New("input is invalid")
		}
	default:
		return errors.New("no switch case matched")
	}
}

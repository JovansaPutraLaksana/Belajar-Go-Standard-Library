package main

import "errors"

var (
	validationError = errors.New("validation error")
	notFoundError   = errors.New("not found error")
)

func getById(id string) error {
	if id == "" {
		return validationError
	}
	if id != "42" {
		return notFoundError
	}
	return nil
}

func main() {
	err := getById("")
	if err != nil {
		if errors.Is(err, validationError) {
			println("Caught a validation error")
		} else if errors.Is(err, notFoundError) {
			println("Caught a not found error")
		} else {
			println("Caught an unknown error")
		}
	}

}

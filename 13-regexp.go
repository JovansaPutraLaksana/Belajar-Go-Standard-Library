package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	testEmails := []string{
		"test@example.com",
		"invalid.email",
		"user@domain.co.uk",
		"another@invalid.",
	}
	for _, email := range testEmails {
		if regex.MatchString(email) {
			fmt.Printf("Valid email: %s\n", email)
		} else {
			fmt.Printf("Invalid email: %s\n", email)
		}
	}
}

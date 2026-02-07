package main

import "fmt"

func main() {
	//%s for strings
	//%d for integers
	//%f for floating-point numbers
	//%t for booleans
	//%v for default format
	//\n for new line

	// Printf allows formatted strings
	name := "Jovan"
	age := 21
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	// Sprintf returns a formatted string
	greeting := fmt.Sprintf("Hello, %s! You are %d years old.", name, age)
	fmt.Println(greeting)
	// Println prints with a newline
	fmt.Println("This is a simple print line.")
}

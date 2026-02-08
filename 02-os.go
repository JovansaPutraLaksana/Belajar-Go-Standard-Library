package main

import (
	"fmt"
	"os"
)

// package os provides a platform-independent interface to operating system functionality.
func main() {
	// go run 02-os.go arg1 arg2 arg3
	// Print command-line arguments
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	// Attempt to read a file that does not exist
	// _, err := os.ReadFile("non_existent_file.txt")
	// if err != nil {
	// 	// Check if the error is of type *os.PathError
	// 	if pathErr, ok := err.(*os.PathError); ok {
	// 		fmt.Println("Caught a PathError:")
	// 		fmt.Printf("Operation: %s\n", pathErr.Op)
	// 		fmt.Printf("Path: %s\n", pathErr.Path)
	// 		fmt.Printf("Error: %s\n", pathErr.Err)
	// 	} else {
	// 		fmt.Println("Caught an unknown error:", err)
	// 	}
	// }

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
	} else {
		fmt.Println("Hostname:", hostname)
	}
}

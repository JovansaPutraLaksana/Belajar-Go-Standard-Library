package main

import (
	"flag"
	"fmt"
)

// package flag implements command-line flag parsing.
// Run the program with flags, for example:
// go run 03-flag.go -host=localhost -port=5433 -user=admin -password=mypassword
func main() {
	host := flag.String("host", "localhost", "Put your database host")
	port := flag.Int("port", 5432, "Put your database port")
	user := flag.String("user", "root", "Put your database user")
	password := flag.String("password", "root", "Put your database password")

	flag.Parse()
	fmt.Printf("Host: %s\n", *host)
	fmt.Printf("Port: %d\n", *port)
	fmt.Printf("User: %s\n", *user)
	fmt.Printf("Password: %s\n", *password)
}

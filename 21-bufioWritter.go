package main

import (
	"bufio"
	"os"
)

func main() {
	writter := bufio.NewWriter(os.Stdout)
	defer writter.Flush()
	writter.WriteString("Hello, World!\n")
	writter.WriteString("This is a buffered writer example.\n")
}

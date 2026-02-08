package main

import (
	"fmt"
	"strings"
)

// package strings implements simple functions to manipulate UTF-8 encoded strings.
func main() {
	// Example string
	str := "  Hello, Go Strings!  "
	fmt.Println("Original string:", str)

	// Trim spaces
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed string:", trimmed)

	// Convert to uppercase
	upper := strings.ToUpper(trimmed)
	fmt.Println("Uppercase string:", upper)

	// Replace substring
	replaced := strings.ReplaceAll(upper, "GO", "GOLANG")
	fmt.Println("Replaced string:", replaced)

	// Split string
	parts := strings.Split(replaced, " ")
	fmt.Println("Split string:", parts)

	// Join strings
	joined := strings.Join(parts, "-")
	fmt.Println("Joined string:", joined)

	// Check if substring is present
	contains := strings.Contains(joined, "GOLANG")
	fmt.Println("Contains 'GOLANG':", contains)

	// Find index of substring
	index := strings.Index(joined, "STRINGS")
	fmt.Println("Index of 'STRINGS':", index)

	// Repeat string
	repeated := strings.Repeat("Go! ", 3)
	fmt.Println("Repeated string:", repeated)

}

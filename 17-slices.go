package main

import (
	"fmt"
	"slices"
)

func main() {
	// Creating and initializing a slice
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Original slice:", numbers)

	// Appending elements to a slice
	numbers = append(numbers, 60, 70)
	fmt.Println("After appending:", numbers)

	// Copying a slice
	copiedNumbers := make([]int, len(numbers))
	slices.Copy(copiedNumbers, numbers)
	fmt.Println("Copied slice:", copiedNumbers)

	// Sorting a slice
	slices.Sort(copiedNumbers)
	fmt.Println("Sorted slice:", copiedNumbers)

	// Searching in a slice
	index := slices.Index(copiedNumbers, 30)
	fmt.Println("Index of 30 in sorted slice:", index)

	// Reversing a slice
	slices.Reverse(copiedNumbers)
	fmt.Println("Reversed slice:", copiedNumbers)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time
	now := time.Now()
	fmt.Println("Current Time:", now)
	// Get specific components
	year, month, day := now.Date()
	hour, min, sec := now.Clock()
	fmt.Printf("Date: %d-%02d-%02d\n", year, month, day)
	fmt.Printf("Time: %02d:%02d:%02d\n", hour, min, sec)
	// Format the time
	formatted := now.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Time:", formatted)
	// Parse a time string
	parsed, err := time.Parse("2006-01-02 15:04:05", "2023-01-01 12:00:00")
	if err != nil {
		fmt.Println("Error parsing time:", err)
	} else {
		fmt.Println("Parsed Time:", parsed)
	}
	// Calculate duration
	duration := now.Sub(parsed)
	fmt.Println("Duration since parsed time:", duration)
}

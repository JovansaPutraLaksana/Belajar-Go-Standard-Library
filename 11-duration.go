package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a duration of 2 hours, 30 minutes, and 45 seconds
	duration := 2*time.Hour + 30*time.Minute + 45*time.Second + 500*time.Millisecond
	fmt.Println("Duration:", duration)
	// Get the total number of seconds in the duration
	totalSeconds := duration.Seconds()
	fmt.Println("Total Seconds:", totalSeconds)
	// Parse a duration string
	parsedDuration, err := time.ParseDuration("1h15m30s")
	if err != nil {
		fmt.Println("Error parsing duration:", err)
	} else {
		fmt.Println("Parsed Duration:", parsedDuration)
	}
	// Add duration to current time
	now := time.Now()
	futureTime := now.Add(duration)
	fmt.Println("Current Time:", now)
	fmt.Println("Future Time after adding duration:", futureTime)
}

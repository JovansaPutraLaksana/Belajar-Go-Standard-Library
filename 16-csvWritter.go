package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	records := [][]string{
		{"name", "age", "city"},
		{"Alice", "30", "New York"},
		{"Bob", "25", "Los Angeles"},
	}
	file, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error creating file:", err.Error())
		return
	}
	writer := csv.NewWriter(file)
	err = writer.WriteAll(records)
	if err != nil {
		fmt.Println("Error writing CSV:", err.Error())
		return
	}
	writer.Flush()
	file.Close()
	fmt.Println("CSV file created successfully.")
}

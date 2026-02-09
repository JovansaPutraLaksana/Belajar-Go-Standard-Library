package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	data := `name,age,city
Alice,30,New York
Bob,25,Los Angeles`
	reader := csv.NewReader(strings.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err.Error())
		return
	}
	for _, record := range records {
		fmt.Println(record)
	}

}

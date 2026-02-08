package main

import (
	"fmt"
	"strconv"
)

// package strconv implements conversions to and from string representations of basic data types.
func main() {
	// Convert string to integer
	intStr := "123"
	intVal, err := strconv.Atoi(intStr)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
	} else {
		fmt.Printf("String to Integer: %s -> %d\n", intStr, intVal)
	}
	// Convert integer to string
	intVal2 := 456
	intStr2 := strconv.Itoa(intVal2)
	fmt.Printf("Integer to String: %d -> %s\n", intVal2, intStr2)
	// Convert string to float
	floatStr := "3.14"
	floatVal, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
	} else {
		fmt.Printf("String to Float: %s -> %f\n", floatStr, floatVal)
	}
	// Convert float to string
	floatVal2 := 2.718
	floatStr2 := strconv.FormatFloat(floatVal2, 'f', 2, 64)
	fmt.Printf("Float to String: %f -> %s\n", floatVal2, floatStr2)
	// Convert string to boolean
	boolStr := "true"
	boolVal, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Println("Error converting string to boolean:", err)
	} else {
		fmt.Printf("String to Boolean: %s -> %t\n", boolStr, boolVal)
	}
	// Convert boolean to string
	boolVal2 := false
	boolStr2 := strconv.FormatBool(boolVal2)
	fmt.Printf("Boolean to String: %t -> %s\n", boolVal2, boolStr2)
}

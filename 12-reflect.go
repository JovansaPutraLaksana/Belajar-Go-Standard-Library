package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"100"`
	Age  int
}

func readField(value any) {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Struct {
		fmt.Println("Expected a struct")
		return
	}
	for i := 0; i < v.NumField(); i++ {
		typeInfo := v.Type().Field(i)
		fmt.Printf("Field %d: %s (%s)\n", i, typeInfo.Name, typeInfo.Type)
		field := v.Field(i)
		fmt.Printf("Field %d: %v\n", i, field.Interface())
	}
}

func readTag(value any) {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Struct {
		fmt.Println("Expected a struct")
		return
	}
	for i := 0; i < v.NumField(); i++ {
		typeInfo := v.Type().Field(i)
		fmt.Printf("Field %d: %s (%s)\n", i, typeInfo.Name, typeInfo.Type)
		tag := typeInfo.Tag.Get("required")
		if tag != "" {
			fmt.Printf("Field %d 'required' tag: %s\n", i, tag)
		}
		tag = typeInfo.Tag.Get("max")
		if tag != "" {
			fmt.Printf("Field %d 'max' tag: %s\n", i, tag)
		}
	}
}

func isValid(value any) {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Struct {
		fmt.Println("Expected a struct")
		return
	}
	for i := 0; i < v.NumField(); i++ {
		typeInfo := v.Type().Field(i)
		tag := typeInfo.Tag.Get("required")
		if tag == "true" {
			field := v.Field(i)
			if !field.IsValid() || field.Interface() == nil || field.Interface() == "" {
				fmt.Printf("Field %d is required but is empty\n", i)
				return
			}
		}
	}
	fmt.Println("All required fields are valid")
}

func main() {
	// var x float64 = 3.456
	// v := reflect.ValueOf(x)
	// fmt.Println("type:", v.Type())
	// fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	// fmt.Println("value:", v.Float())

	s := Sample{Name: "Alice", Age: 30}
	v2 := reflect.ValueOf(s)
	fmt.Println("type:", v2.Type())
	fmt.Println("kind is struct:", v2.Kind() == reflect.Struct)
	fmt.Println("number of fields:", v2.NumField())

	readField(s)
	readTag(s)
	isValid(s)
}

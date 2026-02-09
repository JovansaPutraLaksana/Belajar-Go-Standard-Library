package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.FromSlash("hello/world.go"))
	fmt.Println(filepath.ToSlash("hello\\world.go"))
	fmt.Println(filepath.Split("hello/world.go"))
	fmt.Println(filepath.Clean("hello/../world//main.go"))
	fmt.Println(filepath.VolumeName("C:\\hello\\world.go"))
	fmt.Println(filepath.IsAbs("/hello/world.go"))
	fmt.Println(filepath.IsLocal("/hello/world.go"))
	fmt.Println(filepath.Rel("hello", "hello/world/main.go"))
	fmt.Println(filepath.Match("*.go", "world.go"))
	fmt.Println(filepath.Join("hello", "world", "main.go"))
}

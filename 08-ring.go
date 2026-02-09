package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring with 5 elements
	r := ring.New(5)
	// Populate the ring with values (tambah data)
	for i := 0; i < r.Len(); i++ {
		r.Value = fmt.Sprintf("Element %d", i+1)
		r = r.Next()
	}
	// Iterate through the ring and print elements
	r.Do(func(p interface{}) {
		fmt.Println(p)
	})
}

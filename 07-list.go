package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Create a new list
	myList := list.New()
	// Add elements to the list
	myList.PushBack("First")
	myList.PushBack("Second")
	myList.PushBack("Third")
	// Iterate through the list and print elements
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	var head *list.Element = myList.Front()
	fmt.Println("Head element:", head.Value)
}

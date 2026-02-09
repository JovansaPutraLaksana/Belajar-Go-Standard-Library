package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (u UserSlice) Len() int {
	return len(u)
}
func (u UserSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func main() {
	users := UserSlice{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	sort.Sort(users)

	for _, user := range users {
		fmt.Printf("%s: %d\n", user.Name, user.Age)
	}
}

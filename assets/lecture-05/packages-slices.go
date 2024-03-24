package main

import (
	"cmp"
	"fmt"
	"slices"
)

// START OMIT

type User struct {
	firstname string
	lastname  string
}

func main() {
	users := []User{
		User{
			firstname: "Russ",
			lastname:  "Cox",
		},
	}
	fmt.Println(users)
	users = slices.Insert(users, 0, User{"Rob", "Pike"})
	fmt.Println(users)
	slices.SortFunc(users, func(a, b User) int {
		return cmp.Compare(a.lastname, b.lastname)
	})
	fmt.Println(users)
	users = slices.Delete(users, 0, 2)
	fmt.Println(users)
}

// END OMIT

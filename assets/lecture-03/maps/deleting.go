package main

import (
	"fmt"
)

// START OMIT

// Key represents unique identification of a user.
type Key struct {
	id   uint32
	role string
}

// User represents a system user.
type User struct {
	id      uint32
	name    string
	surname string
}

// MIDDLE OMIT

func main() {
	users := map[Key]User{
		Key{1, "root"} : User{
			id:      1,
			name:    "Linus",
			surname: "Torvalds",
		},
		Key{2, "gopher"}: User{
			id:      2,
			name:    "Rob",
			surname: "Pike",
		},
	}

	fmt.Println(users)
	delete(users, Key{1, "root"}) // existing key
	fmt.Println(users)
	delete(users, Key{1000, "somebody else"}) // non-existent key
	fmt.Println(users)

	clear(users)
	fmt.Println(users)
}

// END OMIT

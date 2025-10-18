package main

import "fmt"

// START OMIT

// Key is a new data type
type Key struct {
	id   uint32
	role string
}

// User is a new data type
type User struct {
	id      uint32
	name    string
	surname string
}

// MIDDLE OMIT

func main() {
	m1 := make(map[Key]User)
	k1 := Key{
		id:   1,
		role: "root",
	}
	m1[k1] = User{
		id:      1,
		name:    "Linus",
		surname: "Torvalds",
	}
	k2 := Key{
		id:   2,
		role: "gopher",
	}
	m1[k2] = User{
		id:      2,
		name:    "Rob",
		surname: "Pike",
	}
	fmt.Println(m1)
}

// END OMIT

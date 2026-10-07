package main

import "fmt"

// START OMIT

type user struct {
	id   uint32
	name string
}

type registeredUser struct {
	user
	email string
}

func main() {
	robert := registeredUser{
		user: user{ // Initialization using the embedded struct
			id:   1,
			name: "Robert Griesemer",
		},
		email: "robert@griesemer.com",
	}
	linus := registeredUser{
		id:    2,
		name:  "Linus Torvalds", // Since Go 1.27, promoted fields can be used as keys directly
		email: "linus@torvalds.com",
	}
	fmt.Println(linus.name, linus.email, robert.name, robert.email)
}

// END OMIT

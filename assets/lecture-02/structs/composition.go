package main

import "fmt"

// START OMIT

type user struct {
	id   uint32
	name string
}

type registeredUser struct {
	user  user
	email string
}

func main() {
	registeredUser := registeredUser{
		user: user{
			id:   1,
			name: "Linus",
		},
		email: "linus@torvalds.com",
	}

	fmt.Println(registeredUser.user.name, registeredUser.email)
}

// END OMIT

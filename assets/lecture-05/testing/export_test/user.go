package models

type User struct {
	name string
}

func createUser(name string) User {
	return User{
		name: name,
	}
}

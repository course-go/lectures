package main

import "fmt"

// User is a new data type
type User struct {
	id      uint32
	name    string
	surname string
}

// START OMIT

func main() {
	u := User{
		id:      1,
		name:    "Linus",
		surname: "Torvalds",
	}
	fmt.Println(u)

	var pName *string
	var pSurname *string
	pName = &u.name
	pSurname = &u.surname
	fmt.Println(pName)
	fmt.Println(pSurname)
	fmt.Println(*pName)
	fmt.Println(*pSurname)

	*pName = "Rob"
	*pSurname = "Pike"
	fmt.Println(*pName)
	fmt.Println(*pSurname)
	fmt.Println(u)
}

// END OMIT

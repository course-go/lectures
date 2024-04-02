package main

import (
	"encoding/json"
	"fmt"
)

// START OMIT

type User struct {
	Name        string
	Active      bool
	Age         uint
	Address     Address
	Phones      []Phone
	AssignedCar *Car
	Children    []string
}

type Address struct {
	Street     string
	City       string
	PostalCode string
}

type Phone struct {
	Type   string
	Number string
}

type Car struct {
	// ...
}

// MIDDLE OMIT

func main() {
	u := User{
		Name:   "Bob",
		Active: true,
		Age:    27,
		Address: Address{
			Street:     "Botanická 68A",
			City:       "Brno - Královo pole",
			PostalCode: "60200",
		},
		Phones: []Phone{
			{
				Type:   "mobile",
				Number: "269 555-1234",
			},
		},
		AssignedCar: nil,
		Children:    []string{"Bob", "Rob"},
	}

	userBytes, _ := json.Marshal(u)
	fmt.Println(string(userBytes))
}

// END OMIT

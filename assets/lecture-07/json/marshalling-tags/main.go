package main

import (
	"encoding/json"
	"fmt"
)

// START OMIT

type User struct {
	Name        string   `json:"name,omitzero"`
	Active      bool     `json:"-"` // Always omitted
	Age         uint     `json:"age,omitzero"`
	Address     Address  `json:"address,omitzero"`
	Phones      []Phone  `json:"phone_numbers,omitempty"`
	AssignedCar *Car     `json:"assigned_car,omitzero"`
	Children    []string `json:"children,omitempty"`
}

type Address struct {
	Street     string `json:"street,omitzero"`
	City       string `json:"city,omitzero"`
	PostalCode string `json:"postal_code,omitzero"`
}

type Phone struct {
	Type   string `json:"type,omitzero"`
	Number string `json:"number,omitzero"`
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
			{Type: "mobile", Number: "269 555-1234"},
		},
		AssignedCar: nil, // Will be omitted
		Children:    []string{"Bob", "Rob"},
	}

	userBytes, _ := json.Marshal(u)
	fmt.Println(string(userBytes))

	var u2 User
	_ = json.Unmarshal(userBytes, &u2)
	fmt.Println(u2)
}

// END OMIT

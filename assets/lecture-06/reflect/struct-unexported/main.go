package main

import (
	"fmt"
	"reflect"
)

// START OMIT

type User struct {
	ID       int
	username string
}

func main() {
	user := User{23, "skidoo"}
	value := reflect.ValueOf(&user).Elem()
	valueType := value.Type()
	fmt.Println("Fields:")
	for i := range value.NumField() {
		field := value.Field(i)
		fieldType := valueType.Field(i)
		fmt.Printf("\t%d: %s %s = %v\n", i, fieldType.Name, field.Type(), field.Interface())
	}
}

// END OMIT

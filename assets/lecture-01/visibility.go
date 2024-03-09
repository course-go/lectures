package main

import (
	"fmt"
	"math"
)

const (
	notExported = "This is not exported!"
	Exported    = "This exported!"
)

func main() {
	fmt.Println(math.Pi) // exported Pi constant
}

package main

import (
	"errors"
	"fmt"
	"os"
)

// START OMIT
var ErrCannotGetNumber = errors.New("cannot get number")

func number() (int, error) {
	return 0, ErrCannotGetNumber
}

func calculate() (int, error) {
	number, err := number()
	if err != nil {
		return 0, fmt.Errorf("cannot calculate: %w", err) // w for wrap
	}
	return number + 5, nil
}

func main() {
	calculation, err := calculate()
	if errors.Is(err, ErrCannotGetNumber) { // perhaps retry?
		return
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

// END OMIT

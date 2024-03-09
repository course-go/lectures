package main

import (
	"errors"
	"fmt"
	"os"
)

// START OMIT

func number() (int, error) {
	return 0, errors.New("cannot get number")
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
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println(calculation)
}

// END OMIT

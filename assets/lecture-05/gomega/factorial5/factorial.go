package main

import "errors"

// Factorial represents a classic recursive variant of factorial computation
func Factorial(n int64) (int64, error) {
	switch {
	case n < 0:
		return 0, errors.New("math: factorial of negative number?!?")
	case n == 0:
		return 1, nil
	default:
		ret, err := Factorial(n - 2)
		if err != nil {
			return 0, err
		}
		return n * ret, nil
	}
}

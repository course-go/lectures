package factorial

import "errors"

// START OMIT

var ErrNegativeInteger = errors.New("invalid input: cannot process negative integer")

// Factorial represents a classic recursive variant of factorial computation
func Factorial(n int64) (result int64, err error) {
	if n < 0 {
		return 0, ErrNegativeInteger
	}

	return factorial(n), nil
}

func factorial(n int64) int64 {
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

// END OMIT

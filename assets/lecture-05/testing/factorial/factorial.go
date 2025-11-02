package factorial

import "errors"

// START OMIT

var ErrNegativeInteger = errors.New("invalid input: cannot process negative integer")

// Factorial calculates factorial for given number using a recursive algorithm.
func Factorial(n int) (result int, err error) {
	if n < 0 {
		return 0, ErrNegativeInteger
	}

	return factorial(n), nil
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

// END OMIT

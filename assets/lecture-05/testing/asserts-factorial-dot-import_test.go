package main

// START OMIT

import (
	"testing"

	. "github.com/stretchr/testify/assert" // Mind the dot!
)

func TestFactorialForZero(t *testing.T) {
	result, err := Factorial(0)
	Nil(t, err)
	Equal(t, result, 1)
}

func TestFactorialForNegative(t *testing.T) {
	_, err := Factorial(-1)
	NotNil(t, err)
	ErrorIs(t, err, NegativeIntegerErr)
}

func TestFactorialForSmallNumber(t *testing.T) {
	result, err := Factorial(5)
	Nil(t, err)
	Greater(t, result, 10)
	Less(t, result, 10000)
}

// END OMIT

package main

// START OMIT

import (
	"testing"

	. "github.com/stretchr/testify/assert" // Mind the dot!
)

func TestFactorialForZero(t *testing.T) {
	result := Factorial(0)
	Equal(t, result, 1)
}

func TestFactorialForOne(t *testing.T) {
	result := Factorial(1)
	Equal(t, result, 1)
}

func TestFactorialForSmallNumber(t *testing.T) {
	result := Factorial(5)
	Greater(t, result, 10)
	Less(t, result, 10000)
}

// END OMIT

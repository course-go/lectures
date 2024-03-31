package main

// START OMIT

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorialForZero(t *testing.T) {
	result := Factorial(0)
	assert.Equal(t, result, 1)
}

func TestFactorialForOne(t *testing.T) {
	result := Factorial(1)
	assert.Equal(t, result, 1)
}

func TestFactorialForSmallNumber(t *testing.T) {
	result := Factorial(5)
	assert.Greater(t, result, 10)
	assert.Less(t, result, 10000)
}

// END OMIT

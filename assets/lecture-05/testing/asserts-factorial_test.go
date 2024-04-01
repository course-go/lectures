package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// START OMIT

func TestFactorialForZero(t *testing.T) {
	result, err := Factorial(0)
	assert.Nil(t, err)
	assert.Equal(t, 1, result)
}

func TestFactorialForOne(t *testing.T) {
	result, err := Factorial(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, result)
}

func TestFactorialForNegative(t *testing.T) {
	_, err := Factorial(-1)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, NegativeIntegerErr)
}

func TestFactorialForSmallNumber(t *testing.T) {
	result, err := Factorial(5)
	assert.Nil(t, err)
	assert.Greater(t, result, 10)
	assert.Less(t, result, 10000)
}

// END OMIT

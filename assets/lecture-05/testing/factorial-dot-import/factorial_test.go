package main

// START OMIT

import (
	"testing"

	"github.com/course-go/lectures/assets/lecture-05/testing/factorial"
	. "github.com/stretchr/testify/assert" // Mind the dot!
)

func TestFactorialOfZero(t *testing.T) {
	result, err := factorial.Factorial(0)
	Nil(t, err)
	Equal(t, 1, result)
}

func TestFactorialOfNegativeNumber(t *testing.T) {
	_, err := factorial.Factorial(-1)
	NotNil(t, err)
	ErrorIs(t, err, factorial.ErrNegativeInteger)
}

func TestFactorialOfSmallNumber(t *testing.T) {
	result, err := factorial.Factorial(5)
	Nil(t, err)
	Greater(t, result, 10)
	Less(t, result, 10000)
}

// END OMIT

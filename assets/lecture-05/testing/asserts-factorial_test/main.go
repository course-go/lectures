package main

import (
	"testing"

	factorial "github.com/course-go/lectures/assets/lecture-05/testing/recursive-factorial"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// START OMIT

func TestFactorialForZero(t *testing.T) {
	result, err := factorial.Factorial(0)
	require.Nil(t, err)
	assert.Equal(t, 1, result)
}

func TestFactorialForOne(t *testing.T) {
	result, err := factorial.Factorial(1)
	require.Nil(t, err)
	assert.Equal(t, 1, result)
}

func TestFactorialForNegative(t *testing.T) {
	_, err := factorial.Factorial(-1)
	require.NotNil(t, err)
	assert.ErrorIs(t, err, factorial.ErrNegativeInteger)
}

func TestFactorialForSmallNumber(t *testing.T) {
	result, err := factorial.Factorial(5)
	require.Nil(t, err)
	assert.Greater(t, result, 10)
	assert.Less(t, result, 10000)
}

// END OMIT

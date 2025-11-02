package factorial_test

import (
	"testing"

	"github.com/course-go/lectures/assets/lecture-05/testing/factorial"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// START OMIT

func TestFactorialOfZero(t *testing.T) {
	result, err := factorial.Factorial(0)
	require.Nil(t, err)
	assert.Equal(t, 1, result)
}

func TestFactorialOfOne(t *testing.T) {
	result, err := factorial.Factorial(1)
	require.Nil(t, err)
	assert.Equal(t, 1, result)
}

func TestFactorialOfNegativeNumber(t *testing.T) {
	_, err := factorial.Factorial(-1)
	require.NotNil(t, err)
	assert.ErrorIs(t, err, factorial.ErrNegativeInteger)
}

func TestFactorialOfSmallNumber(t *testing.T) {
	result, err := factorial.Factorial(5)
	require.Nil(t, err)
	assert.Greater(t, result, 10)
	assert.Less(t, result, 10000)
}

// END OMIT

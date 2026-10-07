package numbers_test

import (
	"math"
	"testing"

	"github.com/course-go/lectures/assets/lecture-05/testing/numbers"
)

// START OMIT

func TestAdd(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"ZeroPlusZero_ReturnsZero", 0, 0, 0},
		{"TwoPlusOne_ReturnsThree", 2, 1, 3},
		{"TwoPlusMinusTwo_ReturnsZero", 2, -2, 0},
		{"MaxIntPlusZero_ReturnsMaxInt", math.MaxInt, 0, math.MaxInt},
		{"MaxIntPlusOne_OverflowsToMinInt", math.MaxInt, 1, math.MinInt},
		{"MaxIntPlusMinInt_ReturnsMinusOne", math.MaxInt, math.MinInt, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := numbers.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("%d + %d should be %d, got %d instead", tt.a, tt.b, tt.expected, result)
			}
		})
	}
}

// END OMIT

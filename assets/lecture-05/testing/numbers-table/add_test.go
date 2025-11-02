package numbers_test

import (
	"math"
	"testing"

	"github.com/course-go/lectures/assets/lecture-05/testing/numbers"
)

// START OMIT

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Add zeroes", 0, 0, 0},
		{"Add positive numbers", 2, 1, 3},
		{"Add mixed numbers", 2, -2, 0},
		{"Add no-op", math.MaxInt32, 0, math.MaxInt32},
		{"Add int32 overflow", math.MaxInt32, 1, math.MinInt32},
		{"Add int32 extremes", math.MaxInt32, math.MinInt32, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numbers.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("%d + %d should be %d, got %d instead", tt.a, tt.b, tt.expected, result)
			}
		})
	}
}

// END OMIT

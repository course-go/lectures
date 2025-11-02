package numbers_test

import (
	"testing"

	numbers "github.com/course-go/lectures/assets/lecture-05/testing/numbers-unexported"
)

func TestAdd(t *testing.T) {
	// `Add` is unreachable as it is not exported.
	result := numbers.Add(1, 2)
	if result != 3 {
		t.Error("1+2 should be 3, got ", result, "instead")
	}
}

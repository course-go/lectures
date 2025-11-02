package numbers_test

import (
	"testing"

	"github.com/course-go/lectures/assets/lecture-05/testing/numbers"
)

func TestAdd(t *testing.T) {
	result := numbers.Add(1, 2)
	if result != 3 {
		t.Error("1+2 should be 3, got:", result, "instead")
	}
}

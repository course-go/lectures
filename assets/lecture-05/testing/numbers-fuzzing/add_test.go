package numbers_test

import (
	"testing"

	"github.com/course-go/lectures/assets/lecture-05/testing/numbers"
)

// START OMIT

func FuzzAdd(f *testing.F) {
	f.Add(1, 6)
	f.Fuzz(func(t *testing.T, a int, b int) {
		t.Logf("Fuzzing: a=%d, b=%d", a, b)
		expected := a + b
		actual := numbers.Add(a, b)
		if actual != expected {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})
}

// END OMIT

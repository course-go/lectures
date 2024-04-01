package numbers_test

import (
	"test/internal/numbers"
	"testing"
)

// START OMIT

func FuzzAdd(f *testing.F) {
	f.Add(1, 6)                               // Seed the fuzz corpus
	f.Fuzz(func(t *testing.T, a int, b int) { // Fuzz target
		t.Logf("Fuzzing: a=%d, b=%d", a, b)
		expected := a + b
		actual := numbers.Add(a, b)
		if actual != expected {
			t.Errorf("expected: %d, actual: %d", expected, actual)
		}
	})
}

// END OMIT

package numbers_test

import (
	"testing"

	"path/to/local/package/numbers"
)

func TestAdd(t *testing.T) {
	result := numbers.Add(1, 2)
	if result != 3 {
		t.Error("1+2 should be 3, got:", result, "instead")
	}
}

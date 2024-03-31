package numbers_test

import (
	"path/to/local/package/numbers"
	"testing"
)

func TestAdd(t *testing.T) {
	result := numbers.Add(1, 2)
	if result != numbers.Add(1, 2) {
		t.Error("1+2 should be 3, got ", result, "instead")
	}
}

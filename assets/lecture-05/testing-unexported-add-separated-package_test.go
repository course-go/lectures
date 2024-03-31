package numbers_test

import (
	"path/to/local/package/numbers"
	"testing"
)

func TestAdd(t *testing.T) {
	result := numbers.Add(1, 2) // Will not work as `add` is not exported
	if result != 3 {
		t.Error("1+2 should be 3, got ", result, "instead")
	}
}

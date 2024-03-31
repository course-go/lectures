package numbers_test

import (
	"path/to/local/package/numbers"
	"testing"
)

// START OMIT

func TestAdd_OneAndTwo_ReturnsThree(t *testing.T) {
	t.Parallel() // Note the call to `Parallel`
	result := numbers.Add(1, 2)
	if result != 3 {
		t.Error("1+2 should be 3, got ", result, "instead")
	}
}

func TestAdd_TenAndTwenty_ReturnsThirty(t *testing.T) {
	t.Parallel() // Note the call to `Parallel`
	result = numbers.Add(10, 20)
	if result != 30 {
		t.Error("10+20 should be 30, got ", result, "instead")
	}
}

func TestAdd_OneAndZero_ReturnsOne(t *testing.T) {
	result := numbers.Add(1, 0)
	if result != 1 {
		t.Error("1+0 should be 1, got ", result, "instead")
	}
}

// END OMIT

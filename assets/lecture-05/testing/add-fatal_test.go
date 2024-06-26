package numbers_test

import "testing"

// START OMIT

func TestAdd(t *testing.T) {
	t.Log("starting test", t.Name())
	result := add(1, 2)
	if result != 3 {
		t.Fatal("1+2 should be 3, got ", result, "instead")
	}

	result = add(10, 20)
	if result != 30 {
		t.Fatal("10+20 should be 30, got ", result, "instead")
	}
}

func TestAddZero(t *testing.T) {
	t.Skip("Skipping test because why not...")
	result := add(1, 0)
	if result != 1 {
		t.Fatal("1+0 should be 1, got ", result, "instead")
	}
}

// END OMIT

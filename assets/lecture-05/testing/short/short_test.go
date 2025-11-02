package short_test

import (
	"testing"
	"time"
)

// START OMIT

func TestLong(t *testing.T) {
	if testing.Short() { // Check if `-short` flag was specified.
		t.SkipNow()
	}
	time.Sleep(1 * time.Hour) // Long test, no doubt...
}

// END OMIT

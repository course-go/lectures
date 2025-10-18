package numbers_test

import (
	"testing"
	"time"
)

// START OMIT

func TestLong(t *testing.T) {
	if t.Short() { // Will not run if `-short` option was set
		t.SkipNow()
	}
	time.Sleep(1 * time.Minute) // Long test, no doubt...
}

// END OMIT

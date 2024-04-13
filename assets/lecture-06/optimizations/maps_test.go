package optimizations_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

// START OMIT

func BenchmarkPreallocatedMapInsert(b *testing.B) {
	m := make(map[uuid.UUID]time.Time, b.N)
	t := time.Now()
	for range b.N {
		b.StopTimer()
		id := uuid.New()
		b.StartTimer()
		m[id] = t
	}
}

func BenchmarkMapInsert(b *testing.B) {
	m := make(map[uuid.UUID]time.Time)
	t := time.Now()
	for range b.N {
		b.StopTimer()
		id := uuid.New()
		b.StartTimer()
		m[id] = t
	}
}

// END OMIT

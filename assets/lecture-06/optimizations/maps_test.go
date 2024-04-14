package optimizations_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

// COMPOUND-KEY-VALUE START OMIT

type value struct{}

type key struct {
	ID      int
	payload [100]byte
}

// COMPOUND START OMIT

func BenchmarkInsertIntoPreallocatedMapCompoundKey(b *testing.B) {
	m := make(map[key]value, b.N)
	for i := 0; i < b.N; i++ {
		k := key{
			ID: i,
		}
		m[k] = value{}
	}
}

func BenchmarkInsertIntoEmptyMapCompoundKey(b *testing.B) {
	m := map[key]value{}
	for i := 0; i < b.N; i++ {
		k := key{
			ID: i,
		}
		m[k] = value{}
	}
}

// COMPOUND END OMIT

// UUID START OMIT

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

// UUID END OMIT

// SET-FILL START OMIT

type ID int

func genID(i int) ID {
	return ID(3*i + 1)
}

func fillInMap(b *testing.B, n int) map[ID]struct{} {
	b.StopTimer()

	m := make(map[ID]struct{}, n)

	for i := 0; i < n; i++ {
		id := genID(i)
		m[id] = struct{}{}
	}

	b.StartTimer()
	return m
}

func fillInSlice(b *testing.B, n int) []ID {
	b.StopTimer()

	s := make([]ID, n)

	for i := 0; i < n; i++ {
		id := genID(i)
		s[i] = id
	}

	b.StartTimer()
	return s
}

// SET-FILL END OMIT
// SET START OMIT

func BenchmarkFindInMap(b *testing.B) {
	s := fillInSlice(b, 1_000_000)
	m := fillInMap(b, 1_000_000)
	performBenchmarkFindInMap(b, m)
	performBenchmarkFindInSlice(b, s)
}

func performBenchmarkFindInMap(b *testing.B, m map[ID]struct{}) {
	items := len(m)
	for i := 0; i < b.N; i++ {
		_, found := m[genID(i%items)]
		if !found {
			b.Fatal("not found")
		}
	}
}

func performBenchmarkFindInSlice(b *testing.B, s []ID) {
	items := len(s)
	for i := 0; i < b.N; i++ {
		found := false
		id := genID(i % items)
		for _, p := range s {
			if p == id {
				found = true
				break
			}
		}
		if !found {
			b.Fatal("not found")
		}
	}
}

// SET END OMIT

package optimizations_test

import (
	"testing"
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
	m := make(map[key]value)
	for i := 0; i < b.N; i++ {
		k := key{
			ID: i,
		}
		m[k] = value{}
	}
}

// COMPOUND END OMIT

// MAP-INSERT START OMIT

func BenchmarkMapInsert(b *testing.B) {
	m := make(map[int]struct{})
	for id := range b.N {
		m[id] = struct{}{}
	}
}

func BenchmarkPreallocatedMapInsert(b *testing.B) {
	m := make(map[int]struct{}, b.N)
	for id := range b.N {
		m[id] = struct{}{}
	}
}

// MAP-INSERT END OMIT

// SET-FILL START OMIT

type ID int

func genID(i int) ID {
	return ID(3*i + 1)
}

func fillInMap(n int) map[ID]struct{} {
	m := make(map[ID]struct{}, n)
	for i := 0; i < n; i++ {
		id := genID(i)
		m[id] = struct{}{}
	}
	return m
}

func fillInSlice(n int) []ID {
	s := make([]ID, n)
	for i := 0; i < n; i++ {
		id := genID(i)
		s[i] = id
	}
	return s
}

// SET-FILL END OMIT

const findInSetSize = 100

// SET START OMIT

func BenchmarkFindInMap(b *testing.B) {
	m := fillInMap(findInSetSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, found := m[genID(i%findInSetSize)]
		if !found {
			b.Fatal("not found")
		}
	}
}

// SET MIDDLE OMIT

func BenchmarkFindInSlice(b *testing.B) {
	s := fillInSlice(findInSetSize)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found := false
		id := genID(i % findInSetSize)
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

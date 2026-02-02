package for_test

import "testing"

const (
	size     = 1_000
	megabyte = 1 << 20
)

// SLICE START OMIT

type Account struct {
	Balance int
	Data    [megabyte]byte
}

func BenchmarkSliceForRangeValue(b *testing.B) {
	accounts := seedAccounts()
	var total int
	for b.Loop() {
		for _, account := range accounts {
			total += account.Balance
		}
	}
}

func BenchmarkSliceForRangeIndex(b *testing.B) {
	accounts := seedAccounts()
	var total int
	for b.Loop() {
		for i := range accounts {
			total += accounts[i].Balance
		}
	}
}

// SLICE END OMIT
// MAP START OMIT

func BenchmarkMapForLoopValue(b *testing.B) {
	m := seedMap()
	var sum int
	for b.Loop() {
		for _, value := range m {
			sum += value
		}
	}
}

func BenchmarkMapForLoopKey(b *testing.B) {
	m := seedMap()
	var sum int
	for b.Loop() {
		for key := range m {
			sum += m[key]
		}
	}
}

// MAP END OMIT

func seedMap() map[int]int {
	m := make(map[int]int, size)
	for i := range size {
		m[i] = i
	}

	return m
}

func seedAccounts() []Account {
	s := make([]Account, size)
	for i := range size {
		s[i] = Account{
			Balance: i,
		}
	}

	return s
}

package loops_test

import "testing"

// START OMIT

func BenchmarkMapForRangeValue(b *testing.B) {
	accounts := seedAccountMap() // map[int]Account
	var total int
	for b.Loop() {
		for _, account := range accounts {
			total += account.Balance
		}
	}
}

func BenchmarkMapForRangeKey(b *testing.B) {
	accounts := seedAccountMap() // map[int]Account
	var total int
	for b.Loop() {
		for id := range accounts {
			total += accounts[id].Balance
		}
	}
}

// END OMIT

func seedAccountMap() map[int]Account {
	const size = 1_000

	accounts := make(map[int]Account, size)
	for i := range size {
		accounts[i] = Account{
			Balance: i,
		}
	}

	return accounts
}

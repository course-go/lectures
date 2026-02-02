package loops_test

import "testing"

// START OMIT

func BenchmarkSliceForRangeValue(b *testing.B) {
	accounts := seedAccountSlice() // []Account
	var total int
	for b.Loop() {
		for _, account := range accounts {
			total += account.Balance
		}
	}
}

func BenchmarkSliceForRangeIndex(b *testing.B) {
	accounts := seedAccountSlice() // []Account
	var total int
	for b.Loop() {
		for i := range accounts {
			total += accounts[i].Balance
		}
	}
}

// END OMIT

func seedAccountSlice() []Account {
	const size = 1_000

	accounts := make([]Account, 0, size)
	for i := range size {
		accounts = append(accounts, Account{
			Balance: i,
		})
	}

	return accounts
}

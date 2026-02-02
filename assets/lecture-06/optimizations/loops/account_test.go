package loops_test

const (
	megabyte = 1 << 20
)

// START OMIT

type Account struct {
	Balance int
	Data    [megabyte]byte
}

// END OMIT

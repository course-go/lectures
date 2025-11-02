package adder

//go:generate go tool mockgen -destination mocks/adder_mock.go -package addermock . Adder

type Adder interface {
	Add(a, b int) int
}

package main

import "fmt"

// START OMIT

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func fileSize() ByteSize {
	return 5 * MB
}

func main() {
	fmt.Printf("%d", int(fileSize()))
}

// END OMIT

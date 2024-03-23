package main

// START OMIT

type UnsignedInteger interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type SignedInteger interface {
	int | int8 | int16 | int32 | int64
}

type Integer interface {
	SignedInteger | UnsignedInteger
}

type Float interface {
	float32 | float64
}

type Complex interface {
	complex64 | complex128
}

type Number interface {
	Integer | Float | Complex
}

// END OMIT

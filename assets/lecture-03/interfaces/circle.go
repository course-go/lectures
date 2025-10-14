package main

// START OMIT

type I1 interface {
	I2
	A()
}

type I2 interface {
	I3
	B()
}

type I3 interface {
	I1
	C()
}

// END OMIT

func main() {}

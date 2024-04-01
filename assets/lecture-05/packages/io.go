package main

// START OMIT

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

// END OMIT

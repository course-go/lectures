package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

// START OMIT

type httpErr struct {
	err  error
	code int
}

func (h httpErr) Error() string {
	return h.err.Error()
}

func (h httpErr) Unwrap() error {
	return h.err
}

func validateRequest(_ http.Request) error {
	return httpErr{err: errors.New("invalid email address"), code: 400}
}

// MIDDLE OMIT

func processRequest(req http.Request) error {
	err := validateRequest(req)
	if err != nil {
		return fmt.Errorf("cannot validate request: %w", err)
	}
	// additional processing...
	return nil
}

func main() {
	err := processRequest(http.Request{})
	if herr, ok := errors.AsType[httpErr](err); ok { // Go 1.26+, previously errors.As
		fmt.Printf("Failed! Code: %d\n", herr.code)
		os.Exit(2)
	}

	if err != nil {
		fmt.Println("Failed!", err)
		os.Exit(2)
	}

	fmt.Println("Success!")
}

// END OMIT

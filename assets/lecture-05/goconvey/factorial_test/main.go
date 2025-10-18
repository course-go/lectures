package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// START OMIT

func TestFactorialForZero(t *testing.T) {
	Convey("0! should be equal 1", t, func() {
		So(Factorial(0), ShouldEqual, 1)
	})
}

func TestFactorialForSmallNumberNegative(t *testing.T) {
	Convey("20! should be between 1 and 10000", t, func() {
		So(Factorial(20), ShouldBeBetween, 1, 10000)
	})
}

func TestFactorialForTen(t *testing.T) {
	Convey("10! should be equal to 3628800", t, func() {
		So(Factorial(10), ShouldEqual, 3628800)
	})
}

func TestFactorialForBigNumber(t *testing.T) {
	Convey("20! should be greater than zero", t, func() {
		So(Factorial(20), ShouldBeGreaterThan, 0)
	})
}

// END OMIT

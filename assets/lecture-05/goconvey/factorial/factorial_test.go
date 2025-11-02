package factorial_test

import (
	"testing"

	"github.com/course-go/lectures/assets/lecture-05/testing/factorial"
	. "github.com/smartystreets/goconvey/convey"
)

// START OMIT

func TestFactorialForZero(t *testing.T) {
	Convey("0! should be equal 1", t, func() {
		output, err := factorial.Factorial(0)
		So(output, ShouldEqual, 1)
		So(err, ShouldBeNil)
	})
}

func TestFactorialForSmallNumber(t *testing.T) {
	Convey("8! should be between 1 and 1 000 000", t, func() {
		output, err := factorial.Factorial(8)
		So(output, ShouldBeBetween, 1, 1_000_000)
		So(err, ShouldBeNil)
	})
}

func TestFactorialForNegativeNumber(t *testing.T) {
	Convey("-1! should return negative integer error", t, func() {
		_, err := factorial.Factorial(-1)
		So(err, ShouldBeError, factorial.ErrNegativeInteger)
	})
}

// END OMIT

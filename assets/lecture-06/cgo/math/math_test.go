package math_test

import (
	"math"
	"testing"

	cmath "github.com/course-go/lectures/assets/lecture-06/cgo/math"
)

const gamma = 0.6628236482

// START OMIT

func BenchmarkGammaGo(b *testing.B) {
	for b.Loop() {
		math.Gamma(gamma)
	}
}

func BenchmarkGammaC(b *testing.B) {
	for b.Loop() {
		cmath.Ctgamma(gamma)
	}
}

// END OMIT

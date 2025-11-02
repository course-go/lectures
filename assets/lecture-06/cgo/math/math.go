package math

// START OMIT

// #cgo LDFLAGS: -lm
//#include <math.h>
import "C"

// Csin is a binding to C sin function.
func Csin(num float64) float64 {
	return float64(C.sin(C.double(num)))
}

// Ccos is a binding to C cos function
func Ccos(num float64) float64 {
	return float64(C.cos(C.double(num)))
}

// Ctgamma is a binding to C tgamma function
func Ctgamma(num float64) float64 {
	return float64(C.tgamma(C.double(num)))
}

// END OMIT

package calculator

import "github.com/course-go/lectures/assets/lecture-05/gomock/adder"

// START OMIT

type Calculator struct {
	adder adder.Adder
}

func New(adder adder.Adder) Calculator {
	return Calculator{
		adder: adder,
	}
}

func (c *Calculator) Add(a, b int) int {
	return c.adder.Add(a, b)
}

// END OMIT

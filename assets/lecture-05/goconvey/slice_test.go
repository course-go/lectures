package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// START OMIT

func TestSliceManipulation(t *testing.T) {
	Convey("Given an empty slice", t, func() {
		s := make([]int, 0)
		Convey("The slice should be empty initially", func() {
			So(s, ShouldBeEmpty)
		})
		Convey("When an item is added", func() {
			s = append(s, 1)
			Convey("The slice should not be empty", func() {
				So(s, ShouldNotBeEmpty)
			})
			Convey("The slice length should be one", func() {
				So(len(s), ShouldEqual, 1)
			})
			Convey("When another item is added", func() {
				s = append(s, 2)
				Convey("The slice length should be two", func() {
					So(len(s), ShouldEqual, 2)
				})
			})
		})
	})
}

// END OMIT

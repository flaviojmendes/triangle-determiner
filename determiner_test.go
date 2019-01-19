package main

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)


func TestDetermineEquilateral(t *testing.T) {
	convey.Convey("Determine that a Triangle where a=5, b=5 and c=5 is Equilateral", t, func() {
		convey.So(determine([]string{"5","5","5"}), convey.ShouldEqual, Equilateral)
	})
}

func TestDetermineIsosceles(t *testing.T) {
	convey.Convey("Determine that a Triangle where a=5.1, b=5.1 and c=4 is Isosceles", t, func() {
		convey.So(determine([]string{"5.1","5.1","10.1"}), convey.ShouldEqual, Isosceles)
	})
}

func TestDetermineScalene(t *testing.T) {
	convey.Convey("Determine that a Triangle where a=5, b=3 and c=4.53 is Scalene", t, func() {
		convey.So(determine([]string{"5","3","4.53"}), convey.ShouldEqual, Scalene)
	})
}

func TestDetermineTriangleInvalidAmountOfSides(t *testing.T) {
	convey.Convey("Determine that a Triangle where a=5, b=3 and does not have c", t, func() {
		convey.So(func(){determine([]string{"5","3"})},convey.ShouldPanicWith, "it is expected an Array with 3 sides of a Triangle")
	})
}


func TestDetermineTriangleInvalidAAndB(t *testing.T) {
	convey.Convey("Determine that a Triangle where a=a, b=b and c=4", t, func() {
		convey.So(func(){determine([]string{"a","b","4"})},convey.ShouldPanicWith, "all sides of a Triangle must be numeric")
	})
}

func TestDetermineTriangleWithFourSides(t *testing.T) {
	convey.Convey("Determine that a Triangle where a=5, b=3, c=4 and has an extra side", t, func() {
		convey.So(func(){determine([]string{"5","3","4","3"})}, convey.ShouldPanicWith, "it is expected an Array with 3 sides of a Triangle")
	})
}

func TestDetermineTriangleNegativeC(t *testing.T) {
	convey.Convey("Determine that a Triangle where a=5, b=3 and c=-1.2", t, func() {
		convey.So(func(){determine([]string{"5","3","-1.2"})}, convey.ShouldPanicWith, "the sum of the lengths of any two sides must be greater than or equal to the length of the remaining side")
	})
}

func TestDetermineTriangleShortAAndB(t *testing.T) {
	convey.Convey("Determine that a Triangle where a=1, b=1 and c=4", t, func() {
		convey.So(func(){determine([]string{"1","1","4"})}, convey.ShouldPanicWith, "the sum of the lengths of any two sides must be greater than or equal to the length of the remaining side")
	})
}

func TestDetermineTriangleShortAAndC(t *testing.T) {
	convey.Convey("Determine that a Triangle where a=1, b=68.3 and c=1", t, func() {
		convey.So(func(){determine([]string{"1","68.3","1"})}, convey.ShouldPanicWith, "the sum of the lengths of any two sides must be greater than or equal to the length of the remaining side")
	})
}


func TestDetermineTriangleWithNotNumericSide(t *testing.T) {
	convey.Convey("Determine that a Triangle where a=1, b=68.3 and c=invalidValue", t, func() {
		convey.So(func(){determine([]string{"1","68.3","invalidValue"})}, convey.ShouldPanicWith, "all sides of a Triangle must be numeric")
	})
}
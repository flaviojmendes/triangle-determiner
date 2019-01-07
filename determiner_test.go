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
		convey.So(determine([]string{"5.1","5.1","4"}), convey.ShouldEqual, Isosceles)
	})
}

func TestDetermineScalene(t *testing.T) {
	convey.Convey("Determine that a Triangle where a=5, b=3 and c=4.53 is Scalene", t, func() {
		convey.So(determine([]string{"5","3","4.53"}), convey.ShouldEqual, Scalene)
	})
}

func TestDetermineTriangleInvalidAmountOfSides(t *testing.T) {
	convey.Convey("Determine that a Triangle where a=5, b=3 and does not have c", t, func() {
		convey.ShouldPanicWith(func(){determine([]string{"5","3"})}, "A Triangle must have 3 sides")
	})
}

func TestDetermineTriangleNegativeC(t *testing.T) {
	convey.Convey("Determine that a Triangle where a=5, b=3 and c=-1.2", t, func() {
		convey.ShouldPanicWith(func(){determine([]string{"5","3","-1.2"})}, "A Triangle must have sides greater than 0")
	})
}

func TestDetermineTriangleShortAAndB(t *testing.T) {
	convey.Convey("Determine that a Triangle where a=1, b=1 and c=4", t, func() {
		convey.ShouldPanicWith(func(){determine([]string{"1","1","4"})}, "A Triangle must have sides greater than 0")
	})
}

func TestDetermineTriangleShortAAndC(t *testing.T) {
	convey.Convey("Determine that a Triangle where a=1, b=68.3 and c=1", t, func() {
		convey.ShouldPanicWith(func(){determine([]string{"1","68.3","1"})}, "A Triangle must have sides greater than 0")
	})
}
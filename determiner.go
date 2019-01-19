package main

import (
	"errors"
	"log"
	"strconv"
)

func determine(givenSides []string) TriangleType {

	// Parsing sides into a numeric and mathematical representation type
	a, b, c := parseSides(givenSides)

	// All sides are different, then it is Scalene
	if a != b && b != c && c != a {
		return Scalene
	// All sides are equal, so it is Equilateral
	} else if a == b && b == c {
		return Equilateral
	// The only option left is Isosceles that has two equal sides.
	} else {
		return Isosceles
	}
}

func parseSides(givenSides []string) (float64,float64,float64) {

	if givenSides == nil || len(givenSides) != 3 {
		log.Panic("it is expected an Array with 3 sides of a Triangle")
	}

	// Parsing the sides to Float64 type
	a,errA := strconv.ParseFloat(givenSides[0],64)
	b,errB := strconv.ParseFloat(givenSides[1],64)
	c,errC := strconv.ParseFloat(givenSides[2],64)

	if errA != nil ||  errB != nil || errC != nil {
		log.Panic("all sides of a Triangle must be numeric")
	}

	validateSides(a,b,c)

	return a,b,c

}

func validateSides(a, b, c float64) {
	// Checking by the Inequality Theorem that it is a Valid Triangle
	inequalityError := errors.New("the sum of the lengths of any two sides must be greater than or equal to the length of the remaining side")

	if a + b <= c {
		log.Panic(inequalityError)
	}

	if b + c <= a {
		log.Panic(inequalityError)
	}

	if c + a <= b {
		log.Panic(inequalityError)
	}
}




package main

import (
	"errors"
	"log"
	"strconv"
)

func determine(givenSides []string) TriangleType {

	// Parsing sides into a numeric and mathematical representation type
	a, b, c, err := parseSides(givenSides)

	if err != nil {
		log.Panic(err)
	}

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

func parseSides(givenSides []string) (float64,float64,float64, error) {

	if givenSides == nil || len(givenSides) != 3 {
		return 0, 0, 0, errors.New("it is expected an Array with 3 sides of a Triangle")
	}

	// Parsing the sides to Float64 type
	a,err := strconv.ParseFloat(givenSides[0],64)
	b,err := strconv.ParseFloat(givenSides[1],64)
	c,err := strconv.ParseFloat(givenSides[2],64)

	if err != nil {
		return 0, 0, 0, errors.New("all sides of a Triangle must be numeric")
	}

	_,err = hasValidSides(a,b,c)

	return a,b,c,err

}

func hasValidSides(a, b, c float64) (bool, error) {
	// Checking by the Inequality Theorem that it is a Valid Triangle
	inequalityError := errors.New("the sum of the lengths of any two sides must be greater than or equal to the length of the remaining side")

	if a + b <= c {
		return false, inequalityError
	}

	if b + c <= a {
		return false, inequalityError
	}

	if c + a <= b {
		return false, inequalityError
	}

	// If got here it is a valid Triangle!
	return true, nil
}




// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
    // Pick values for the following identifiers used by the test program.
    NaT Kind = iota // not a triangle
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	arr := [3]float64{a, b, c}

	for _, num := range arr {
		// Not a triangle
		if num <= 0 || math.IsNaN(num) || math.IsInf(num, 0) || math.IsInf(num, -1) {
			return NaT
		}
	}

	// Not a triangle
	if a + b < c || c + a < b || b + c < a {
		return NaT
	}

	// Three side have same length
	if a == b && a == c {
		return Equ
	}

	// Triangle have at least two side same length
	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}

// Package triangle ...
package triangle

import "math"

// Kind is the type of triangle
type Kind int

// Types of triangles
const (
	NaT Kind = iota // Not a triangle
	Equ             // Equilateral
	Iso             // Isosceles
	Sca             // Scalene
)

// KindFromSides returns the type of triangle, given three sides
func KindFromSides(a, b, c float64) Kind {
	if a+b < c || b+c < a || c+a < b ||
		IsInvalidSide(a) || IsInvalidSide(b) || IsInvalidSide(c) {
		return NaT
	}
	if a == b && a == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}

// IsInvalidSide returns true if the side is 0, NaN, or Inf
func IsInvalidSide(side float64) bool {
	if side == 0 || math.IsInf(side, 0) || math.IsNaN(side) {
		return true
	}
	return false
}

// Package triangle ...
package triangle

import "math"

// Kind is the type of triangle
type Kind string

const (
	// NaT is not a triangle
	NaT = "NaT"
	// Equ is an quilateral triangle
	Equ = "Equ"
	// Iso is an isosceles triangle
	Iso = "Iso"
	// Sca is a scalene triangle
	Sca = "Sca"
)

// KindFromSides ...
func KindFromSides(a, b, c float64) Kind {
	if a == math.Inf(1) || b == math.Inf(1) || a == math.NaN() ||
		b == math.NaN() || a == math.Inf(-1) || b == math.Inf(-1) {
		return NaT
	}
	if a+b >= c && b+c >= a && c+a >= b && a+b+c != 0 {
		if a == b && a == c {
			return Equ
		}
		if a == b || a == c || b == c {
			return Iso
		}
		return Sca
	}

	return NaT
}

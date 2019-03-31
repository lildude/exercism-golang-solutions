//Package triangle gives triangle sides length and tells if a triangle is equilateral, isosceles, or scalene.
package triangle

import "math"

// Kind represents the kind of triangle.
type Kind string

// Triangle kind constants
const (
	NaT Kind = "NaT"
	Equ Kind = "Equ"
	Iso Kind = "Iso"
	Sca Kind = "Sca"
)

func isTriangle(a, b, c float64) bool {
	return isSide(a) && isSide(b) && isSide(c) && a+b >= c && a+c >= b && b+c >= a
}

func isSide(s float64) bool {
	return s > 0 && !math.IsNaN(s) && !math.IsInf(s, 1)
}

func isEqu(a, b, c float64) bool {
	return a == b && b == c && a == c
}

func isIso(a, b, c float64) bool {
	return a == b || b == c || a == c
}

func isSca(a, b, c float64) bool {
	return a != b && b != c && a != c
}

// KindFromSides returns the kind of triangle based on the values for the sides passed to it
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	} else if isEqu(a, b, c) {
		return Equ
	} else if isIso(a, b, c) {
		return Iso
	} else if isSca(a, b, c) {
		return Sca
	}
	return NaT
}

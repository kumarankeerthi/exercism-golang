// excercism Triage
package triangle

import (
	"math"
)

type Kind int

const (
	
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// given three sides of a triangle this function will return the Kind of Triange.
func KindFromSides(a, b, c float64) Kind {

	var k Kind

	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) ||
		math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) ||
		math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		a <= 0 || b <= 0 || c <= 0 || (a+b < c) || (a+c < b) || (b+b < a) {
		k = Kind(NaT)
	} else if a == b && b == c && c == a {
		k = Kind(Equ)
	} else if a == b || b == c || c == a {
		k = Kind(Iso)
	} else if a != b && b != c && c != a {
		k = Kind(Sca)
	}
	return k
}

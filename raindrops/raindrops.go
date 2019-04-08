// Package raindrops
package raindrops

import (
	"strconv"
)

// Convert returns a string that is ia combination of "Pling" , "Plang" , "Plong" based on the factors of a given number
func Convert(n int) (result string) {
	if n%3 == 0 {
		result += "Pling"
	}
	if n%5 == 0 {
		result += "Plang"
	}
	if n%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		result = strconv.Itoa(n)
	}
	return

}

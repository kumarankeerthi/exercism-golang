// Package raindrops
package raindrops

import (
	"strconv"
)

// Convert returns a string that is ia combination of "Pling" , "Plang" , "Plong" based on the factors of a given number
func Convert(n int) string {
	var result string
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			if i == 3 {
				result += "Pling"
			} else if i == 5 {
				result += "Plang"
			} else if i == 7 {
				result += "Plong"
			}
		}
	}
	if result == "" {
		result = strconv.Itoa(n)
	}
	return result

}

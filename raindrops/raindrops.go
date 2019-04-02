// Package raindrops
package raindrops

import (
	"strconv"
)

// Convert returns a string that is ia combination of "Pling" , "Plang" , "Plong" based on the factors of a given number
func Convert(n int) string {
	var result string
	factors := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	for _, v := range factors {
		if v == 3 {
			result += "Pling"
		} else if v == 5 {
			result += "Plang"
		} else if v == 7 {
			result += "Plong"
		}
	}
	if result == "" {
		result = strconv.Itoa(n)
	}
	return result

}

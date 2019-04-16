// package scrabble
package scrabble

import (
	"unicode"
)

var weights = map[rune]int{
	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'S': 1, 'T': 1, 'R': 1,
	'D': 2, 'G': 2,
	'B': 3, 'C': 3, 'M': 3, 'P': 3,
	'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
	'K': 5,
	'J': 8, 'X': 8,
	'Q': 10, 'Z': 10,
}

// Score returns the score for a given string based on the
// weightage assinged to each charater and the no.of time it repeats
func Score(s string) (value int) {
	for _, v := range s {
		value += weights[unicode.ToUpper(v)]
	}
	return
}

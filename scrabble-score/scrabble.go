// package scrabble
package scrabble

import (
	"strings"
)

type weight struct {
	Letter string
	Value  int
}

var weights = []weight{
	{"AEIOULNRST", 1},
	{"DG", 2},
	{"BCMP", 3},
	{"FHVWY", 4},
	{"K", 5},
	{"JX", 8},
	{"QZ", 10},
}

// Score returns the score for a given string based on the
// weightage assinged to each charater and the no.of time it repeats
func Score(s string) (value int) {
	value = 0
	for _, v := range s {
		for _, word := range weights {
			if strings.Contains(word.Letter, strings.ToUpper(string(v))) {
				value += word.Value
				break
			}
		}
	}
	return
}

//Package isogram
package isogram

import (
	"unicode"
)

// function check is a given string is isoggram - returns bool
func IsIsogram(s string) bool {
	isIsogram := true
	for i, rune := range s {
		if unicode.IsLetter(rune) {
			for j, rune2 := range s {
				if i != j && unicode.ToLower(rune) == unicode.ToLower(rune2) {
					isIsogram = false
					break
				}
			}
		}
	}
	return isIsogram
}

// BenchmarkIsIsogram-4       59406             19241 ns/op               0 B/op          0 allocs/op

/*
func IsIsogram(s string) bool {
	isIsogram := true
	for i, rune := range s {
		if unicode.IsLetter(rune) {
			for j := i + 1; j < len(s); j++ {
				if string(unicode.ToLower(rune)) == strings.ToLower(string(s[j])) {
					isIsogram = false
					break
				}
			}
		}
	}
	return isIsogram
}

  BenchmarkIsIsogram-4       45283             23541 ns/op              18 B/op         18 allocs/op
*/

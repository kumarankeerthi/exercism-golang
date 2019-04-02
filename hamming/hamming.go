// Package Hamming
package hamming

import (
	"errors"
)

// Distance returns the hamming distance between two DNA strands (string)
func Distance(a, b string) (int, error) {
	hammingDistance := 0
	if len(a) != len(b) {
		return 0, errors.New("false")
	}
	for i, _ := range a {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}

// package dna
package dna

import (
	"errors"
)

type Histogram map[rune]int

type DNA string

// Counts calculate the no.of times a nucleotide occurs in a given DNA strand
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, nucleotide := range d {
		_, ok := h[nucleotide]
		if !ok {
			return h, errors.New("invalid nucleotide in the given DNA strand")
		}
		h[nucleotide]++
	}
	return h, nil
}

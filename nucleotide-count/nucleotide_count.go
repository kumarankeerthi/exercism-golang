// package dna
package dna

import (
	"errors"
	"strings"
)

type Histogram map[rune]int

type DNA string

const validNucleotide = "ACGT"

// Counts calculate the no.of times a nucleotide occurs in a given DNA strand
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	if d == "" {
		return h, nil
	}
	for _, nucleotide := range d {
		if !(strings.Contains(validNucleotide, string(nucleotide))) {
			return h, errors.New("invalid nucleotide in the given DNA strand")
		}
		h[nucleotide] += 1
	}
	return h, nil
}

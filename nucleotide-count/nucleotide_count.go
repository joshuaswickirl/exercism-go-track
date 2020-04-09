package dna

import (
	"errors"
	"regexp"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

var r = regexp.MustCompile("^[ACGT]+$")

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	if d != "" && !r.MatchString(string(d)) {
		return Histogram{}, errors.New("string with invalid nucleotides")
	}
	for _, s := range d {
		h[rune(s)]++
	}
	return h, nil
}

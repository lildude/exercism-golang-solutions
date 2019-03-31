// Package dna provides a method for counting nucleotides in a DNA string
package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, l := range d {
		_, ok := h[l]
		if ok {
			h[l]++
		} else {
			return h, errors.New("Invalid nucleotide")
		}
	}
	return h, nil
}

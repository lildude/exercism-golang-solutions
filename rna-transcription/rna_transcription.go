// Package strand takes a DNA string and returns its RNA complement
package strand

import (
	"strings"
)

// DnaToRnaMap maps DNA nucleotides to their RNA complements
var DnaToRnaMap = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA takes a DNA string and returns its RNA complement
func ToRNA(dna string) string {
	var rna string
	if !strings.ContainsAny(dna, "GCTA") {
		return ""
	}
	for _, l := range dna {
		rna += string(DnaToRnaMap[l])
	}
	return rna
}

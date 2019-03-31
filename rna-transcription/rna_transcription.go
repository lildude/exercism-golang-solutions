// Package strand takes a DNA string and returns its RNA complement
package strand

// DnaToRnaMap maps DNA nucleotides to their RNA complements
var DnaToRnaMap = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA takes a DNA string and returns its RNA complement
func ToRNA(dna string) string {
	rna := ""
	for _, l := range dna {
		rna += string(DnaToRnaMap[l])
	}
	return rna
}

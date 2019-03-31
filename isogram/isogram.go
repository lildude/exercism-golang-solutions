// Package isogram provides a method to determine is a word is an isogram or not
package isogram

import "unicode"

// IsIsogram determines if a word is an isogram or not
func IsIsogram(word string) bool {
	seen := make(map[rune]struct{})
	for _, letter := range word {
		lcLetter := unicode.ToLower(letter)
		if lcLetter == '-' || lcLetter == ' ' {
			continue
		}
		_, yes := seen[lcLetter]
		if yes {
			return false
		}
		seen[lcLetter] = struct{}{}
	}
	return true
}

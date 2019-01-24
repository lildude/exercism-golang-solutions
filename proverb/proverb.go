// Package proverb provides methods to return a proverb based on the words provided.
package proverb

import "fmt"

// Proverb returns a slice of lines for a proverb based o the words provided.
func Proverb(rhyme []string) []string {
	out := []string{}

	for i := range rhyme {
		if i == 0 {
			continue
		}
		out = append(out, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
	}

	if len(rhyme) >= 1 {
		out = append(out, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}

	return out
}

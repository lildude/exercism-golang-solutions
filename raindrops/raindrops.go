// Package raindrops provides a method to convert a number to string based on it's factoring.
package raindrops

import "strconv"

// Convert a number to a string, the contents of which depend on the number's factors.
func Convert(n int) string {
	out := ""
	if n%3 == 0 {
		out += "Pling"
	}
	if n%5 == 0 {
		out += "Plang"
	}
	if n%7 == 0 {
		out += "Plong"
	}

	if out == "" {
		out = strconv.Itoa(n)
	}
	return out
}

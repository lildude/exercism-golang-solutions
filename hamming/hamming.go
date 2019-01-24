package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("String lengths don't match")
	}

	aChars := []rune(a)
	bChars := []rune(b)
	dist := 0

	for i, v := range aChars {
		if v != bChars[i] {
			dist++
		}
	}

	return dist, nil
}

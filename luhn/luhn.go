// Package luhn validates a number using the Luhn algorithm
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid determines if a string is valid as per Luhn's algorithm
func Valid(input string) bool {
	squashed := strings.Replace(input, " ", "", -1)

	if len(squashed) <= 1 {
		return false
	}

	numbers := strings.Split(squashed, "")
	sum := int(0)

	for _, char := range squashed {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	for i := len(numbers) - 1; i >= 0; i-- {
		v, _ := strconv.Atoi(numbers[i])
		n := int(v)
		if (len(numbers)-i)%2 == 0 {
			n = v * 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
	}

	return sum%10 == 0
}

// Package romannumerals provides a method for converting numbers to roman numerals.
package romannumerals

import (
	"errors"
	"sort"
	"strings"
)

//
var romanNumeralMap = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

// ToRomanNumeral converts a normal number to Roman Numerals
func ToRomanNumeral(number int) (string, error) {
	if number <= 0 || number > 3000 {
		return "", errors.New("Only positive integers less than 3000 can be converted to roman numeral")
	}

	roman := ""
	remainder := number

	// Go iterates over maps in a random order so we need to force order by pulling out the keys,
	// ordering them and then iterating through these to pull in the correct element from the map.
	keys := make([]int, 0)
	for k := range romanNumeralMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for remainder > 0 {
		for _, k := range keys {
			if (remainder / k) > 0 {
				roman += strings.Repeat(romanNumeralMap[k], (remainder / k))
			}
			remainder = remainder % k
		}
	}
	return roman, nil
}

// Package romannumerals provides a method for converting numbers to roman numerals.
package romannumerals

import (
	"errors"
)

var romanNumerals = []struct {
	number int
	roman  string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeral converts a normal number to Roman Numerals
func ToRomanNumeral(number int) (string, error) {
	if number <= 0 || number > 3000 {
		return "", errors.New("Only positive integers less than 3000 can be converted to roman numeral")
	}

	roman := ""
	for _, transform := range romanNumerals {
		for number >= transform.number {
			roman += transform.roman
			number -= transform.number
		}
	}
	return roman, nil
}

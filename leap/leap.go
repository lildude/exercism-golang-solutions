// Package leap implements a method for determining if a given year is a leap year.
package leap

// IsLeapYear determines if the given year is a leap year.
func IsLeapYear(y int) bool {
	if y%4 == 0 && (y%100 != 0 || y%400 == 0) {
		return true
	}
	return false
}

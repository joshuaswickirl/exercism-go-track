// Package leap includes a function that calculates if a given year
// is a leap year.
package leap

// IsLeapYear returns true if a given year is a leap year, false if not.
func IsLeapYear(year int) bool {
	if year%4 != 0 || (year%100 == 0 && year%400 != 0) {
		return false
	}
	return true
}

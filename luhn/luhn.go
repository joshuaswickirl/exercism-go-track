package luhn

import (
	"strings"
)

// Valid checks if a number passes Luhn validation
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) <= 1 {
		return false
	}
	numbers := []rune(input)
	var sum int
	everyOther := len(numbers)%2 == 0
	for _, value := range numbers {
		value := int(value - '0')
		if value < 0 || value > 9 {
			return false
		}
		if everyOther {
			value *= 2
			if value > 9 {
				value -= 9
			}
		}
		sum += value
		everyOther = !everyOther
	}
	return sum%10 == 0
}

package luhn

import (
	"strings"
)

// Valid checks if a credit card number is valid
func Valid(cardNumberInput string) bool {
	cardNumber := strings.ReplaceAll(cardNumberInput, " ", "")
	if len(cardNumber) <= 1 {
		return false
	}
	cardNumberSlice := []rune(cardNumber)
	var sum int
	var everyOther bool
	for i := len(cardNumberSlice) - 1; i >= 0; i-- {
		value := int(cardNumberSlice[i] - '0')
		if value < 0 || value > 9 {
			return false
		}
		if everyOther {
			value = value * 2
			if value > 9 {
				value = value - 9
			}
		}
		sum += value
		everyOther = !everyOther
	}
	return sum%10 == 0
}

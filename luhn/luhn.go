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
	numIndexes := len(cardNumberSlice) - 1
	for i := numIndexes; i >= 0; i-- {

		value := int(cardNumberSlice[i] - '0')

		if value < 0 || value > 9 {
			return false
		}

		if (numIndexes == i) || (numIndexes-i)%2 == 0 {
			sum += value
		} else {

			newValue := value * 2

			if newValue > 9 {
				newValue = newValue - 9
			}

			sum += newValue
		}
	}

	return sum%10 == 0
}

package luhn

import (
	"strconv"
	"strings"
)

// Valid checks if a credit card number is valid
func Valid(cardNumberInput string) bool {
	cardNumber := strings.ReplaceAll(cardNumberInput, " ", "")

	// Validate input
	_, err := strconv.Atoi(cardNumber)
	if err != nil || len(cardNumber) <= 1 {
		return false
	}

	cardNumberSlice := []rune(cardNumber)

	var sum int
	numIndexes := len(cardNumberSlice) - 1
	for i := numIndexes; i >= 0; i-- {

		if (numIndexes == i) || (numIndexes-i)%2 == 0 {
			sum += int(cardNumberSlice[i] - '0')
		} else {

			value := int(cardNumberSlice[i] - '0')
			newValue := value * 2

			if newValue > 9 {
				newValue = newValue - 9
			}

			sum += newValue
		}
	}

	if sum%10 != 0 {
		return false
	}

	return true

}

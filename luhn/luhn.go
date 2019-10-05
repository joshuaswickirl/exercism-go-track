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

	cardNumberArray := strings.Split(cardNumber, "")

	var sum int
	numIndexes := len(cardNumberArray) - 1
	for i := numIndexes; i >= 0; i-- {

		if (numIndexes == i) || (numIndexes-i)%2 == 0 {
		} else {

			value, _ := strconv.Atoi(cardNumberArray[i])
			newValue := value * 2

			if newValue > 9 {
				newValue = newValue - 9
			}

			cardNumberArray[i] = strconv.Itoa(newValue)
		}

		value, _ := strconv.Atoi(cardNumberArray[i])
		sum += value

	}

	if sum%10 != 0 {
		return false
	}

	return true

}

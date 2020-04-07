package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode returns a Crypto Square encoded message for a given string.
func Encode(message string) string {
	normalMessage := normalize(message)
	if len(normalMessage) <= 1 {
		return normalMessage
	}
	rectangle, width, height := formRectangle(normalMessage)
	paddingNeeded := int(math.Abs(float64((width * height) - len(normalMessage))))
	encodedMessage := encodeMessage(rectangle, width, height, paddingNeeded)
	return encodedMessage
}

func normalize(message string) string {
	var normalized string
	for _, r := range strings.ToLower(message) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			normalized += string(r)
		}
	}
	return normalized
}

func formRectangle(normalizedMessage string) ([]string, int, int) {
	var rectangle []string
	var numCols, currentCol int
	var newRow string
	// columns >= rows and columns - rows <= 1
	numCols = int(math.Ceil(math.Sqrt(float64(len(normalizedMessage)))))
	for i, r := range normalizedMessage {
		newRow += string(r)
		if currentCol == numCols-1 || i == len(normalizedMessage)-1 {
			rectangle = append(rectangle, newRow)
			newRow = ""
			currentCol = 0
			continue
		}
		currentCol++
	}
	return rectangle, len(rectangle[0]), len(rectangle)
}

func encodeMessage(rectangle []string, width, height, padding int) string {
	var encodedMessage string
	for col := 0; col < width; col++ {
		for row := 0; row < height; row++ {
			if col > width-1-int(padding) && row == height-1 {
				encodedMessage += " "
			} else {
				encodedMessage += string(rectangle[row][col])
			}
			if row == height-1 && col != width-1 {
				encodedMessage += " "
			}
		}
	}
	return encodedMessage
}

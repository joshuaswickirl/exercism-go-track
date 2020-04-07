package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode returns a Crypto Square encoded message for a given string.
func Encode(message string) string {
	normalMessage := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return -1
	}, strings.ToLower(message))
	messageLength := len(normalMessage)
	width := int(math.Ceil(math.Sqrt(float64(messageLength))))
	area := width * (width - 1)
	padding := 0
	if area < messageLength {
		padding = area - messageLength + width
	} else if area > messageLength {
		padding = area - messageLength
	}
	columns := make([]string, width)
	for i, r := range normalMessage + strings.Repeat(" ", padding) {
		columns[i%width] += string(r)
	}
	return strings.Join(columns, " ")
}

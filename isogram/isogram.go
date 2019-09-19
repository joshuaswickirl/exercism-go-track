package isogram

import "unicode"

// IsIsogram returns true if the word passed is an isogram
func IsIsogram(word string) bool {

	var lettersUsed = map[rune]bool{}

	isIsogram := true

	for _, char := range word {

		if char == '-' || char == ' ' {
			continue
		}

		letter := unicode.ToLower(char)
		_, exists := lettersUsed[letter]

		if exists {
			isIsogram = false
			break
		} else {
			lettersUsed[letter] = true
		}
	}

	return isIsogram
}

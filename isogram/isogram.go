package isogram

import "unicode"

// IsIsogram returns true if the word passed is an isogram
func IsIsogram(word string) bool {

	var lettersUsed = map[rune]bool{}

	for _, char := range word {

		if char == '-' || char == ' ' {
			continue
		}

		letter := unicode.ToLower(char)

		if lettersUsed[letter] {
			return false
		}

		lettersUsed[letter] = true
	}

	return true
}

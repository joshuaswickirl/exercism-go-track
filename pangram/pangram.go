package pangram

import (
	"strings"
	"unicode"
)

// IsPangram returns true if a sentence uses all letters in the alphabet.
func IsPangram(sentence string) bool {
	usedLetters := map[rune]bool{}
	for _, char := range strings.ToLower(sentence) {
		if unicode.IsLetter(char) {
			usedLetters[char] = true
		}
	}
	return len(usedLetters) == 26
}

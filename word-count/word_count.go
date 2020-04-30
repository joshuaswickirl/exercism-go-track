package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	f := Frequency{}
	words := strings.FieldsFunc(phrase, func(r rune) bool {
		return r == ' ' || r == ','
	})
	for _, w := range words {
		w = strings.TrimFunc(w, func(r rune) bool {
			return !(unicode.IsLetter(r) || unicode.IsDigit(r))
		})
		w = strings.ToLower(w)
		if w != "" {
			f[w]++
		}
	}
	return f
}

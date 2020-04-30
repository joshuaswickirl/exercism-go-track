package anagram

import (
	"reflect"
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var matches []string
	sLetters := strings.Split(strings.ToLower(subject), "")
	sort.Strings(sLetters)
	for _, candidate := range candidates {
		if strings.ToLower(candidate) == strings.ToLower(subject) {
			continue
		}
		cLetters := strings.Split(strings.ToLower(candidate), "")
		sort.Strings(cLetters)
		if reflect.DeepEqual(cLetters, sLetters) {
			matches = append(matches, candidate)
		}
	}
	return matches
}

package etl

import "strings"

// Transform returns the shiny new format from the legacy data format
func Transform(oldMap map[int][]string) map[string]int {
	output := map[string]int{}
	for point, letters := range oldMap {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = point
		}
	}
	return output
}

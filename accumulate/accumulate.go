package accumulate

// Accumulate returns a slice of strings modified by a given function
func Accumulate(input []string, converter func(string) string) []string {
	var result []string
	for _, s := range input {
		result = append(result, converter(s))
	}
	return result
}

package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	var substrings []string
	for i := 0; i < len(s)-n+1; i++ {
		substrings = append(substrings, s[i:i+n])
	}
	return substrings
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

// First returns the first substring of s with length n.
func First(n int, s string) (first string, ok bool) {
	if len(s) < n {
		return "", false
	}
	return s[0:n], true
}

// Package hamming contains a function that calculates Hamming Distance
package hamming

import "fmt"

// Distance returns the Hamming Distance between two DNA strands
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, fmt.Errorf("%s and %s are of different length", a, b)
	}

	hammingDistance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}

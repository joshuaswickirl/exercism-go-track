// Package hamming contains a function that calculates Hamming Distance
package hamming

import "fmt"

// Distance returns the Hamming Distance between two DNA strands
func Distance(a, b string) (int, error) {

	if len([]rune(a)) != len([]rune(b)) {
		return 0, fmt.Errorf("%s and %s are of different length", a, b)
	}

	hammingDistance := 0
	for i := 0; i < len([]rune(a)); i++ {
		if []rune(a)[i] != []rune(b)[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}

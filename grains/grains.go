package grains

import (
	"errors"
	"math"
)

// Square returns the number of grains on a square
func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("there are 64 squares on a chessboard")
	}
	return uint64(math.Pow(2, float64(input-1))), nil
}

// Total returns the total number of grains on the board
func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		grains, _ := Square(i)
		total += grains
	}
	return total
}

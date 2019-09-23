package diffsquares

import "math"

// SquareOfSum calculates the square of the sum of
func SquareOfSum(n int) int {
	sum := (n*n + n) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of the squares
func SumOfSquares(n int) int {
	var square int
	for i := 0; i <= n; i++ {
		square += i * i
	}
	return square
}

// Difference contrasts the return value of SquareOfSum and SumOfSquares
func Difference(n int) int {
	return int(math.Abs(float64(SquareOfSum(n) - SumOfSquares(n))))
}

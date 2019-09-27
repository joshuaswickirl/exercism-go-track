package diffsquares

import "math"

// SquareOfSum calculates the square of the sum of
func SquareOfSum(n int) int {
	sum := (n*n + n) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of the squares
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference contrasts the return value of SquareOfSum and SumOfSquares
func Difference(n int) int {
	return int(math.Abs(float64(SquareOfSum(n) - SumOfSquares(n))))
}

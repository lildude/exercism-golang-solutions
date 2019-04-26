package diffsquares

// SquareOfSum calculates the square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	t := (n + 1) * n * 1 / 2
	return t * t
}

// SumOfSquares calculates the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference calculates the difference between the square of the sum of the
// first n natural numbers and the sum of the squares of the first n natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

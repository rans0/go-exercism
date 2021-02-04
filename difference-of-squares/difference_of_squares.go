package diffsquares

// test version 2

func SquareOfSum(n int) int {
	ret := (n * (n + 1)) / 2
	return ret * ret
}

func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
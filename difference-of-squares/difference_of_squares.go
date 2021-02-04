package diffsquares

func SquareOfSum(n int) (sum int) {
	for i:=1; ; i++ {
		if i > n {
			return sum * sum
		}
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) (sum int){
	for i:=1; ; i++ {
		if i > n {
			return sum
		}
		sum += i * i
	}
	return sum
}

func Difference(n int) (diff int) {
	sqOfSum := SquareOfSum(n)
	sumOfSq := SumOfSquares(n)

	return sqOfSum - sumOfSq
}
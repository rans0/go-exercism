package prime

func Nth(n int) (prime int, ok bool) {
	var temp []int

	if n < 1 {
		return 0, false
	}

	for i:=2; ; i++ {
		if checkPrime(i, temp) {
			temp = append(temp, i)
			if len(temp) == n {
				return i, true
			}
		}
	}
}

func checkPrime(num int, temp []int) bool {
	for _, value := range temp {
		if num % value == 0 {
			return false
		}
	}
	return true
}
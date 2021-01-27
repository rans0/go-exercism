package prime

func Factors(num int64) []int64 {
	arrays := make([]int64, 0)
	var i int64
	for i=2;i <= num; i++ {
		for num % i == 0 {
			arrays = append(arrays, i)
			num /= i
		}
	}
	return arrays
}
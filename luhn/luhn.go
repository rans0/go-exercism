package luhn

import (
	"strings"
)

// Valid performs Luhn validation on the given number
func Valid(num string) bool {
	num = strings.Replace(num, " ", "", -1)
	if len(num) <= 1 {
		return false
	}

	var sum int32
	for idx, n := range num {
		if n < 48 || n > 57 {
			return false
		}

		temp := num[len(num)-1-idx] - '0'
		if idx % 2 == 1 {
			temp = temp * 2 // if temp == minus, must be X 2 for positive value
			if temp > 9 {
				temp -= 9
			}
		}
		sum += int32(temp)
	}

	return sum % 10 == 0
}
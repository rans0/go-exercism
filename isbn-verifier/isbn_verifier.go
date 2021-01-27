package isbn

import (
	"unicode"
)

func IsValidISBN (isbn string) bool {
	maxLength, checksum := 0, 0
	for _, num := range isbn {
		switch {
		case unicode.IsDigit(num):
			checksum += (10 - maxLength) * int(num - 48) // 48 is equal to '0' zero
			maxLength++
		case num == 'X' && maxLength == 9:
			checksum += 10
			maxLength++
		}
	}
	return maxLength == 10 && checksum % 11 == 0
}
package pangram

import "strings"

func IsPangram (s string) bool {
	lower := strings.ToLower(s)
	for i:='a'; i <= 'z'; i++ {
		// return false if lower is not equal to i rune
		// alphabet not equal to ascii
		if !(strings.ContainsRune(lower, i)) {
			return false
		}
	}
	return true
}

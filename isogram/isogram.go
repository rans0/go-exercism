package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(s string) bool {
	s = strings.ToLower(s)
	for idx, char := range s {
		if unicode.IsLetter(char) && strings.ContainsRune(s[idx+1:], char) {
			return false
		}
	}
	return true
}
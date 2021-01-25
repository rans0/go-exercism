package reverse

import "unicode/utf8"

func Reverse (s string) string{
	length := len(s)
	buffer := make([]byte, length)
	for i:=0; i < length; {
		r, index := utf8.DecodeRuneInString(s[i:])
			i += index
			utf8.EncodeRune(buffer[length-i:], r)
	}
	return string(buffer)
}
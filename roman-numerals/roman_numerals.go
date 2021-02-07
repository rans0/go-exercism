package romannumerals

import (
	"errors"
)

/*
Algorithm :
input 1024 as arabic

loop 1 :
idx = 0, arabic >= arabics[idx] => True
arabics[idx] = 1000
arabic = 1024 - 1000 = 24
roman = romans[idx]
roman = M

loop 2 :
idx = 1, arabic >= arabics[idx] => False

loop 9 :
arabic = 24
idx 8, arabic >= arabics[idx] => True
arabic = 24 - 10 = 14
roman = romans[idx]
roman = X

loop 10 :
arabic = 14
idx 9, arabic >= arabics[idx] => True
arabic = 14 - 9 = 5
roman = romans[idx]
roman = XI

loop 11 :
arabic = 5
idx 10, arabic >= arabics[idx] => True
arabic = 5 - 5 = 0
roman = romans[idx]
roman = V

loop... END :
roman = XXIV
*/

var arabics = []int{
	1000, 900, 500, 400,
	100, 90, 50, 40,
	10, 9, 5, 4, 1,
}
var romans = []string{
	"M", "CM", "D", "CD",
	"C", "XC", "L", "XL",
	"X", "IX", "V", "IV", "I",
}

func ToRomanNumeral (arabic int) (roman string, err error) {
	if arabic < 1 || arabic > 3000 { // errror
		return "", errors.New("error")
	}

	for idx, _ := range arabics { // convert to roman
		for arabic >= arabics[idx] {
			arabic -= arabics[idx]
			roman += romans[idx]
		}
	}

	return roman, nil
}
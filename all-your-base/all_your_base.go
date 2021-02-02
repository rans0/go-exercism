package allyourbase

import (
	"errors"
)

func ConvertToBase (base int, lists []int, outputBase int) (result []int, err error) {

	if base < 2 {
		return nil, errors.New("input base must be >= 2")
	}

	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	if len(lists) == 0 {
		lists = append(lists, 0)
	}

	var decimal int
	for _, num := range lists {
		if num < 0 || num >= base {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		decimal = decimal * base + num
	}

	if decimal == 0 {
		result = append(result, 0)
	}

	for decimal > 0 {
		dec := decimal % outputBase
		result = append(result, 0)
		copy(result[1:], result[0:])
		result[0] = dec
		decimal = decimal / outputBase
	}

	return result, nil
}
package collatzconjecture

import "errors"
func CollatzConjecture (input int) (expected int, err error){
	if input <= 0 {
		return 0, errors.New("invalid")
	}

	for true {
		// if input = 1
		if input == 1{
			break
		}
		// even
		if input%2 == 0 {
			input = input / 2
		} else if input % 2 == 1 { // odd
			input = input*3 + 1
		}
		expected++
	}
	return expected, nil
}
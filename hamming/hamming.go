package hamming

import "errors"

func Distance(a, b string) (int, error) {
	count := 0

	// check length string
	if len(a) != len(b) {
		return 0, errors.New("invalid")
	}

	// check char is equal
	for i, _ := range a {
		if a[i] != b[i] {
			count++
		}
	}

	return  count, nil
}

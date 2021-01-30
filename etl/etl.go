package etl

import "strings"

func Transform(input map[int][]string) (output map[string]int) {
	output = make(map[string]int, len(input))
	for key, values := range input {
		for _, value := range values {
			output[strings.ToLower(value)] = key
		}
	}
	return
}
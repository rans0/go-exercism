package scrabble

import (
	"strings"
)

var letterScores = make(map[rune]int)

var class = map[string]int {
	"aeioulnrst": 1,
	"dg": 2,
	"bcmp": 3,
	"fhvwy": 4,
	"k": 5,
	"jx": 8,
	"qz": 10,
}

func init() {
	for letters, score := range class {
		for _, letter := range letters {
			letterScores[letter] = score
		}
	}
}

func Score (string string) (score int) {
	for _, char := range strings.ToLower(string) {
		score += letterScores[char]
	}
	return
}
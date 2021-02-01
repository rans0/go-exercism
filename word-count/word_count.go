package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(sentence string) (freq Frequency) {
	freq = map[string]int{}
	rule := regexp.MustCompile("[a-zA-Z0-9]+(['][a-zA-Z0-9]+)?")

	for _, word := range rule.FindAllString(sentence, -1) {
		word = strings.ToLower(word)
		freq[word]++
	}
	return
}
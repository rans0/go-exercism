package anagram

import (
	"sort"
	"strings"
)

func sorting(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func Detect(subject string, candidates []string) (expected []string) {
	subject = strings.ToLower(subject)
	wordSorted := sorting(subject)

	for _, c := range candidates {
		word := strings.ToLower(c)
		if subject == word || wordSorted != sorting(word){
			continue
		}
		expected = append(expected, c)
	}
	return
}
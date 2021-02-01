package acronym

import "strings"

func Abbreviate(s string) (acronym string) {
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Replace(s, "_", " ", -1)

	words := strings.Fields(s)

	for idx, _ := range words {
		acronym += strings.ToUpper(string(words[idx][0]))
	}

	return
}

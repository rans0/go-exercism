package accumulate

func Accumulate(words []string, convert func(string) string) []string {
	buffer := make([]string, len(words))
	for idx, word := range words {
		buffer[idx] = convert(word)
	}
	return buffer
}

package twofer

func ShareWith(name string) string {
	var sentence string
	if name != "" {
		sentence = "One for " + name + ", one for me."
	} else {
		return "One for you, one for me."
	}
	return sentence
}

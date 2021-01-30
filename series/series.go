package series

func All(n int, s string) (result []string) {
	for i:=0; i < len(s) - n; i++ {
		result = append(result, s[i:i+n])
	}
	return
}

func UnsafeFirst(n int, s string) (result string) {
	if (len(s) < n) {
		return s
	}
	return s[0:n]
}
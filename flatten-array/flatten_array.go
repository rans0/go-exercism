package flatten

func Flatten (lists interface{}) []interface{} {
	result := []interface{}{}
	for _, ok := range lists.([]interface{}) {
		switch value := ok.(type) {
		case []interface{}:
			result = append(result, Flatten(value)...)
		case int:
			result = append(result, value)
		}
	}
	return result
}
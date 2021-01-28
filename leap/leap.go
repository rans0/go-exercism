package leap

func IsLeapYear(number int) bool {
	// n every year that is evenly divisible by 4
	// except every year that is evenly divisible by 100
	// unless the year is also evenly divisible by 400
	return number % 4 == 0 && number % 100 != 0 || number % 400 == 0
}

package armstrong

import (
	"math"
	"strconv"
	"strings"
)

func convertToIntArray(num int) (arrayOfInts []float64){
	letter := strconv.Itoa(num)
	arrayOfStrs := strings.Split(letter, "")

	for _, i := range arrayOfStrs {
		toInt, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		arrayOfInts = append(arrayOfInts, float64(toInt))
	}

	return arrayOfInts
}

func IsNumber (num int) bool {
	var temp float64
	arr := convertToIntArray(num)

	for i:=0; i < len(arr); i++ {
		temp += math.Pow(arr[i], float64(len(arr)))
		if int(temp) == num {
			return true
		}
	}

	return false
}
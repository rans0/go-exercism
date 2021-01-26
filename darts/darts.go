package darts

import "math"

func Score(dart float64, target float64) int{
	var point int
	calculate := math.Sqrt(math.Pow(dart-0, 2) + math.Pow((target-0), 2))
	switch {
	case calculate <= 1 :
		point = 10
	case calculate <= 5 :
		point = 5
	case calculate <= 10 :
		point = 1
	default:
		point = 0
	}
	return point
}
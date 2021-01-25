package space

type Planet string

var secondsInOneYearEarth = 31556952.0

var periodOrbital = map[Planet] float64 {
	"Mercury": 0.2408467,
	"Venus": 0.61519726,
	"Earth": 1,
	"Mars" : 1.8808158,
	"Jupiter": 11.862615,
	"Saturn": 29.447498,
	"Uranus": 84.016846,
	"Neptune": 164.79132,
}

func convertSecondInOneYearPlanet(planet Planet) float64 {
	return secondsInOneYearEarth * periodOrbital[planet]
}

func Age(second float64, planet Planet) float64 {
	return Round( second/convertSecondInOneYearPlanet(planet))
}

func Round(second float64) float64 {
	return float64(int64(second/0.01+0.5)) * 0.01
}

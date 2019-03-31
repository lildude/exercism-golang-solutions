// Package space implements methods for calculating ages based on planet years.
package space

// Planet represents the planet on which we want to calculate the age
type Planet string

const earthSeconds = float64(31557600)

var planetMap = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age calculates the age relative to a planet rotation cycle
func Age(ageInSeconds float64, planet Planet) float64 {
	return ageInSeconds / (planetMap[planet] * earthSeconds)
}

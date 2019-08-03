package space

type Planet string

var orbitConversion = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(ageSeconds float64, planet Planet) (planetAge float64) {
	earthAge := ageSeconds / 60.0 / 60.0 / 24.0 / 365.25
	return earthAge / orbitConversion[planet]
}

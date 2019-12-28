// Exercism side exercise about Space Age.
package space

type Planet string

// Age calculates how old someone would be on any planet.
func Age(f float64, planet Planet) float64 {

	var earthYearInSecs float64 = 31557600

	switch planet{
	case "Earth":
		return f / earthYearInSecs
	case "Mercury":
		return f / (earthYearInSecs * 0.2408467)
	case "Venus":
		return f / (earthYearInSecs * 0.61519726)
	case "Mars":
		return f / (earthYearInSecs * 1.8808158)
	case "Jupiter":
		return f / (earthYearInSecs * 11.862615)
	case "Saturn":
		return f / (earthYearInSecs * 29.447498)
	case "Uranus":
		return f / (earthYearInSecs * 84.016846)
	case "Neptune":
		return f / (earthYearInSecs * 164.79132)
	default:
		return 0
	}

}
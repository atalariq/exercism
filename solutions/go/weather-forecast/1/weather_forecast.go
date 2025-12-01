// Package weather is a program that can forecast the current weather condition of various locations.
package weather

// CurrentCondition represent weather condition.
// CurrentLocation represent location.
var (
	CurrentCondition string
	CurrentLocation  string
)

// Forecast takes city location and condition and return string that report the weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

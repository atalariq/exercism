// Package weather provides the forecast of various locations.
package weather

// CurrentCondition describe current weather condition.
var CurrentCondition string

// CurrentLocation is a location to forecast.
var CurrentLocation string

// Forecast gives the report of weather forecast for given location & condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

//Package weather that tells you the weather.
package weather

//CurrentCondition of the weather.
var CurrentCondition string
//CurrentLocation where user is searching.
var CurrentLocation string

//Forecast function that does the heavy lifting.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

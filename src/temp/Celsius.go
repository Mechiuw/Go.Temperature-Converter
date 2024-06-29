package temp

type Celsius struct {
	Temperature int
	Category    string
}

func ToFahrenheit(celsius Celsius) int {
	return (celsius.Temperature * 9 / 5) + 32
}

package temp

type Fahrenheit struct {
	Temperature int
	Category    string
}

func ToKelvin(f Fahrenheit) float64 {
	return (float64)(f.Temperature-32)*5/9 + 273.15
}

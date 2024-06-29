package temp

type Reaumur struct {
	Temperature int
	Category    string
}

func ToCelsius(r Reaumur) float64 {
	return (float64(r.Temperature) * 5 / 4)
}

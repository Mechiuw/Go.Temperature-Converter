package temp

type Rankine struct {
	Temperature int
	Category    string
}

func ToReaumur(r Rankine) float64 {
	return (float64(r.Temperature) - 491.67) * 4 / 9
}

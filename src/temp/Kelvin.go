package temp

type Kelvin struct {
	Temperature int
	Category    string
}

func ToRankine(k Kelvin) float64 {
	return (float64(k.Temperature) * 1.8)
}

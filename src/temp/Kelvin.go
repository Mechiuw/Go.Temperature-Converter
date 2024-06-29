package temp

type Kelvin struct {
	Temperature int
	Category    int
}

func ToRankine(k Kelvin) float64 {
	return (float64(k.Temperature) * 1.8)
}

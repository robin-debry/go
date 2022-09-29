package star

func Star(s Perimeter) float64 {
	return (s.x + s.y) * 5
}

type Perimeter struct {
	x float64
	y float64
}

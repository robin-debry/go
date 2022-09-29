package star

func Star(s Perimeter) float64 {
	return (s.x + s.y) * 5
}

func StarArea(a Area) float64 {
	return (a.x * a.y) * 10
}

type Perimeter struct {
	x float64
	y float64
}

type Area struct {
	x float64
	y float64
}

package geom

type Star struct {
	length float64
}

func (s Star) Perimeter() float64 {
	return (s.length + s.length) * 5
}

func (s Star) Area() float64 {
	return (s.length * s.length) * 10
}

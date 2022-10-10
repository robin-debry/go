package geom

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Star struct {
	length float64
}

func (s Star) Perimeter() float64 {
	return (s.length + s.length) * 5
}

func (s Star) Area() float64 {
	return (s.length * s.length) * 10
}

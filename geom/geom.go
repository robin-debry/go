package geom

import (
	"math"
)

func PerimeterRectangle(rect Rectangle) float64 {
	return 20
}

type Rectangle struct {
	height float64
	width  float64
}

// Perimeter calculates the perimeter of a width and a height given of a rectangle.
func (r Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}

// Area calculates the area of a width and a height given of a rectangle.
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

type Pentagone struct {
	edge     float64
	apotheme float64
}

func (p Pentagone) Perimeter() float64 {
	return 5 * p.edge
}

func (p Pentagone) Area() float64 {
	return 0.5 * p.Perimeter()* p.apotheme
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

func Perimeter(s Shape) float64 {
	return s.Perimeter()
}


func (h Hexagone) Area() float64 {
	return (3 * math.Sqrt(3) * h.width * h.width) / 2
}

func (h Hexagone) Perimeter() float64 {
	return h.width * 6
}

type Hexagone struct {
	width float64
}

func Area(s Shape) float64 {
	return s.Area()
}


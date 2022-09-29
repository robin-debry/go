package geom

import "math"

// Area calculates the area of a width and a height given of a rectangle.
func Area(width, height float64) float64 {
	return width * height
}

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

type Circle struct {
	radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Shape interface {
	Perimeter() float64
}

func Perimeter(s Shape) float64 {
	return s.Perimeter()
}

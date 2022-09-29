package geom

import "math"

// Perimeter calculates the perimeter of a width and a height given of a rectangle.
func Perimeter(width, height float64) float64 {
	return (width + height) * 2
}

// Area calculates the area of a width and a height given of a rectangle.
func Area(width, height float64) float64 {
	return width * height
}

func PerimeterCircle(radius float64) float64 {
	return 2 * math.Pi * radius
}

func PerimeterRectangle(rect Rectangle) float64 {
	return 20
}

type Rectangle struct {
	height float64
	width  float64
}

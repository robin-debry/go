package geom

// Perimeter calculates the perimeter of a width and a height given of a rectangle.
func Perimeter(width, height int) int {
	return (width + height) * 2
}

// Area calculates the area of a width and a height given of a rectangle.
func Area(width, height int) int {
	return width * height
}

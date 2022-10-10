package geom

import "testing"


func TestPerimeter(t *testing.T) {
	cases := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"perimeter 5 ", Pentagone{5, 3}, 25},
		{"perimeter 2.5 ", Circle{2.5}, 15.707963267948966},
		{"perimeter 4 and 6", Rectangle{4, 6}, 20},
		{"perimeter 7 and 9", Rectangle{7, 9}, 32},
	}

	for i := 0; i < len(cases); i++ {
		c := cases[i]
		t.Run(c.name, func(t *testing.T) {
			got := c.shape.Perimeter()
			if got != c.want {
				t.Fatalf("got %f, want %f", got, c.want)
			}
		})
	}
}


func TestArea(t *testing.T) {
	cases := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Area 5 and 3", Pentagone{5, 3}, 37.5},
		{"Area 4 and 6", Rectangle{4, 6}, 24},
		{"Area 7 and 9", Rectangle{7, 9}, 63},
	}

	for i := 0; i < len(cases); i++ {
		c := cases[i]
		t.Run(c.name, func(t *testing.T) {
			got := c.shape.Area()
			if got != c.want {
				t.Fatalf("got %f, want %f", got, c.want)
			}
		})
	}
}
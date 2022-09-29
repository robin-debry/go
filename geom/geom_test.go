package geom

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("perimeter 4 and 6", func(t *testing.T) {
		got := perimeter(4, 6)
		if got != 20 {
			t.Error("perimeter 4 and 6 is not 20")
		}
	})
	t.Run("perimeter 7 and 9", func(t *testing.T) {
		got := perimeter(7, 9)
		if got != 32 {
			t.Error("perimeter 7 and 9 is not 32")
		}
	})
}
func TestArea(t *testing.T) {
	t.Run("Area 3 and 5", func(t *testing.T) {
		got := Area(3, 5)
		if got != 15 {
			t.Error("Area 3 and 5 is not 15")
		}
	})
	t.Run("Area 3 and 4", func(t *testing.T) {
		got := Area(3, 4)
		if got != 12 {
			t.Error("Area 3 and 4 is not 12")
		}
	})
}

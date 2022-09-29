package star

import "testing"

func TestStar_Perimeter(t *testing.T) {
	t.Run("perimeter 4 and 4", func(t *testing.T) {
		r := Perimeter{
			5, 5,
		}
		got := Star(r)
		if got != 40 {
			t.Error("perimeter 5 and 5 is not 50")
		}
	})
}

func TestStar_Area(t *testing.T) {
	t.Run("area 4 and 4 and 2", func(t *testing.T) {
		a := Area{
			2, 4.58257569495584,
		}
		got := StarArea(a)
		if got != 91.6515138991168 {
			t.Error("Area 2 and 4.58257569495584 and 10 is not 91.6515138991168")
		}
	})
}

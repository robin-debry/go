package geom

import "testing"

func TestStar_Perimeter(t *testing.T) {
	t.Run("perimeter 4 and 4", func(t *testing.T) {
		s := Star{
			length: 5,
		}
		got := s.Perimeter()
		if got != 50 {
			t.Error("perimeter 5 and 5 is not 50")
		}
	})
}

func TestStar_Area(t *testing.T) {
	t.Run("area 4 and 4 and 2", func(t *testing.T) {
		s := Star{
			length: 5,
		}
		got := s.Area()
		if got != 250 {
			t.Error("Area 2 and 4.58257569495584 and 10 is not 250")
		}
	})
}

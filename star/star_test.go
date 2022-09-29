package star

import "testing"

func TestStar_Perimeter(t *testing.T) {
	t.Run("perimeter 4 and 4", func(t *testing.T) {
		r := Perimeter{
			4, 4,
		}
		got := Star(r)
		if got != 40 {
			t.Error("perimeter 4 and 4 is not 40")
		}
	})
}

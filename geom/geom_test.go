package geom

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("perimeter 4 and 6", func(t *testing.T) {
		got := Perimeter(Rectangle{4, 6})
		if got != 20 {
			t.Error("perimeter 4 and 6 is not 20")
		}
	})
	t.Run("perimeter 7 and 9", func(t *testing.T) {
		got := Perimeter(Rectangle{7, 9})
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

func TestCircle_Perimeter(t *testing.T) {
	c := Circle{
		2.5,
	}
	c2 := Circle{
		5.1,
	}
	got := c.Perimeter()
	if got != 15.707963267948966 {
		t.Error("not the good result, got :", got)
	}
	got2 := c2.Perimeter()
	if got2 != 32.044245066615886 {
		t.Error("not the good result, got :", got2)
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	t.Run("perimeter 4 and 6", func(t *testing.T) {
		r := Rectangle{
			4, 6,
		}
		got := r.Perimeter()
		if got != 20 {
			t.Error("perimeter 4 and 6 is not 20")
		}
	})
	t.Run("perimeter 7 and 9", func(t *testing.T) {
		r := Rectangle{
			7, 9,
		}
		got := r.Perimeter()
		if got != 32 {
			t.Error("perimeter 7 and 9 is not 32")
		}
	})
}

func TestPentagone_Perimeter(t *testing.T) {
	t.Run("perimeter 5 ", func(t *testing.T) {
		p := Pentagone{
			5,
		}
		got := p.Perimeter()
		if got != 25 {
			t.Error("perimeter 5 is not 25")
		}
	})
}
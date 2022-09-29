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
func TestCircle_Perimeter(t *testing.T) {
	t.Run("Area 2.5 and 5.1", func(t *testing.T) {
		c := Circle{
			2.5,
		}
		got := c.Perimeter()
		if got != 15.707963267948966 {
			t.Error("not the good result, got :", got)
		}
	})
	t.Run("Area 2.5 and 5.1", func(t *testing.T) {
		c := Circle{
			5.1,
		}
		got := c.Perimeter()
		if got != 32.044245066615886 {
			t.Error("not the good result, got :", got)

		}

	})
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

func TestHexagone_Area(t *testing.T) {
	h := Hexagone{
		4,
	}
	h2 := Hexagone{
		12,
	}
	got := h.Area()
	if got != 41.569219381653056 {
		t.Error("not the good result, got :", got)
	}
	got2 := h2.Area()
	if got2 != 374.1229744348775 {
		t.Error("not the good result, got :", got2)
	}
}

func TestHexagone_Perimeter(t *testing.T) {
	h := Hexagone{
		4,
	}
	h2 := Hexagone{
		12,
	}
	got := h.Perimeter()
	if got != 24 {
		t.Error("not the good result, got :", got)
	}
	got2 := h2.Perimeter()
	if got2 != 72 {
		t.Error("not the good result, got :", got2)
	}
}
func TestPentagone_Perimeter(t *testing.T) {
	t.Run("perimeter 5 ", func(t *testing.T) {
		p := Pentagone{
			5, 3,
		}
		got := p.Perimeter()
		if got != 25 {
			t.Error("perimeter 5 is not 25")
		}
	})
}

func TestPentagoneArea(t *testing.T) {
	t.Run("Area 5 ", func(t *testing.T) {
		p := Pentagone{
			5, 3,
		}
		got := p.Area()
		if got != 37.5 {
			t.Error("Area 5 is not 37.5")
		}
	})

}

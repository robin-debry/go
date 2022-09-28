package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("1+2+3=6", func(t *testing.T) {
		got := sum([]int{
			1, 2, 3,
		})
		if got != 6 {
			t.Error(" it's not 6")
		}
	})
	t.Run("1+2=3", func(t *testing.T) {
		got := sum([]int{
			1, 2,
		})
		if got != 3 {
			t.Error(" it's not 3")
		}
	})
}

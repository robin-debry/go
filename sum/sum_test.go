package sum

import "testing"

func TestSum(t *testing.T) {
	got := sum([]int{
		1, 2, 3,
	})
	if got == 6 {
		return
	} else {
		t.Error(" it's not 6")
	}
}

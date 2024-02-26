package calc

import (
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum(1, 2)
	want := 3
	if got != want {
		t.Errorf("Sum with 1,2 == %d, want %d", got, want)
	}
}

package stack

import "testing"

func TestIsEmpty(t *testing.T) {
	got := new(AStack).IsEmpty()
	if got == false {
		t.Error("s = []. s.IsEmpty() want true, got", got)
	}
}

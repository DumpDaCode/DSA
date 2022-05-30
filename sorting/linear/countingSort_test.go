package linear

import (
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {
	t.Run("Basic Input", func(t *testing.T) {
		arr := []int{0, 0, 1, 1, 0, 0, 1, 0}
		want := []int{0, 0, 0, 0, 0, 1, 1, 1}
		got := CountingSort(arr, 3)
		t.Log(got)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})
}

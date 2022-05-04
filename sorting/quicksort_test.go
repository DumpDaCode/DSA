package sorting

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Run("array is not Sorted", func(t *testing.T) {
		arr := []int{1, 3, 2, 4}
		want := []int{1, 2, 3, 4}
		QuickSort(arr, 0, len(arr)-1)
		got := arr
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})
	t.Run("array contains repeated elements", func(t *testing.T) {
		arr := []int{1, 3, 2, 1, 9, 2}
		want := []int{1, 1, 2, 2, 3, 9}
		QuickSort(arr, 0, len(arr)-1)
		got := arr
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})
	t.Run("array is already sorted", func(t *testing.T) {
		arr := []int{1, 2, 3, 4}
		want := []int{1, 2, 3, 4}
		QuickSort(arr, 0, len(arr)-1)
		got := arr
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})
}

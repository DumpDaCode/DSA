package sorting

import "github.com/DumpDaCode/DSA/constants"

// Good for short input sizes.

func InsertionSort(arr []int) {
	n := len(arr)
	for j := 1; j < n; j++ {
		key := arr[j]
		var i int
		for i = j - 1; i >= 0 && key < arr[i]; i-- {
			arr[i+1] = arr[i]
		}
		arr[i+1] = key
	}
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// Sorts in ascending order
func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		warr := arr[i:]
		min := constants.MaxInt
		key := 0
		for i := 0; i < len(warr); i++ {
			if min > warr[i] {
				min = warr[i]
				key = i
			}
		}
		arr[i], arr[key+i] = arr[key+i], arr[i]
	}
}

package sorting

import (
	"math/rand"
)

func QuickSort(arr []int, start, end int) {
	if start < end {
		pivot := Partition(arr, start, end)
		QuickSort(arr, start, pivot-1)
		QuickSort(arr, pivot+1, end)
	}
}

func Partition(arr []int, start, end int) int {
	pivot := arr[end]
	i := start
	for j := start; j < len(arr)-1; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i += 1
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}

func RandomizedPartition(arr []int, start, end int) int {
	i := rand.Intn(end-start) + start
	arr[i], arr[end-1] = arr[end-1], arr[i]
	return Partition(arr, start, end)
}

func RandomizedQuickSort(arr []int, start, end int) {
	if start < end {
		pivot := RandomizedPartition(arr, start, end)
		QuickSort(arr, start, pivot-1)
		QuickSort(arr, pivot+1, end)
	}
}

func HoarePartion(arr []int, start, end int) int {
	x := arr[start]
	i := start - 1
	j := end + 1
	for {
		for arr[j] >= x {
			j = j - 1
		}
		for arr[i] <= x {
			i = i + 1
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			break
		}
	}
	return j
}

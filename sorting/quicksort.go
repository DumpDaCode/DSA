package sorting

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

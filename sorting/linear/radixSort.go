package linear

func RadixSort(arr []int, k int) {
	for i := 0; i < k; i++ {
		CountingSort(arr, 9)
	}
}

package linear

func CountingSort(arr []int, k int) []int {
	fArr := make([]int, len(arr))
	tArr := make([]int, k)
	for i := 0; i < len(arr); i++ {
		tArr[arr[i]-1] += 1
	}
	for i := 1; i < k; i++ {
		tArr[i] += tArr[i-1]
	}
	for i := len(fArr) - 1; i >= 0; i-- {
		fArr[tArr[arr[i]-1]-1] = arr[i]
		tArr[arr[i]-1] -= 1
	}
	return fArr
}

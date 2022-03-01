package sorting

// Good for large input sizes
// Divide and conquer approach
func Merge(arr []int, p int, q int, r int) {
	left := make([]int, q-p+1)
	copy(left, arr[p:q+1])
	right := make([]int, r-q)
	copy(right, arr[q+1:r+1])
	i := 0
	j := 0
	k := p
	for ; k <= r && i < len(left) && j < len(right); k++ {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
	for ; i < len(left); i++ {
		arr[k] = left[i]
		k++
	}
	for ; j < len(right); j++ {
		arr[k] = right[j]
		k++
	}
}

func MergeSort(arr []int, p int, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(arr, p, q)
		MergeSort(arr, q+1, r)
		Merge(arr, p, q, r)
	}
}

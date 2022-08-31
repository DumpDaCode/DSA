package linear

// CountingSort algorithm is desinged for arrays that contains digits in a pre determined range
func CountingSort(arr []int, k int) []int {
	sortedArr := make([]int, len(arr))
	tempArr := make([]int, k)
	// Counting the number of occurence of each element
	for i := 0; i < len(arr); i++ {
		tempArr[arr[i]-1] += 1
	}
	// Now add the occurence of each element with with count of previous element
	// This results into each element having a number that represents it's correct index
	for i := 1; i < k; i++ {
		tempArr[i] += tempArr[i-1]
	}
	// Now take the index of element from tempArr and place it into correct position in the final array.
	// Also decrement the index of the same element in tempArr.
	// Here we scan the orignal array from right to left.
	for i := len(sortedArr) - 1; i >= 0; i-- {
		sortedArr[tempArr[arr[i]-1]-1] = arr[i]
		tempArr[arr[i]-1] -= 1
	}
	return sortedArr
}

package algorithms

import (
	"github.com/DumpDaCode/DSA/constants"
)

// Holds result.
type resultTuple struct {
	start int
	end   int
	sum   int
}

// Max across the midpoint
func MaxCrossSubArray(arr []int, low int, mid int, high int) resultTuple {
	var res resultTuple = resultTuple{start: 0, end: 0, sum: 0}
	leftSum, rightSum := constants.MinInt, constants.MinInt
	for i := mid; i >= low; i-- {
		res.sum += arr[i]
		if res.sum > leftSum {
			leftSum = res.sum
			res.start = i
		}
	}
	res.sum = 0
	for i := mid + 1; i <= high; i++ {
		res.sum += arr[i]
		if res.sum > rightSum {
			rightSum = res.sum
			res.end = i
		}
	}
	res.sum = leftSum + rightSum
	return res
}

// Divide and conquer approach
func MaxSubArray(arr []int, low int, high int) resultTuple {
	if high == low {
		return resultTuple{start: low, end: high, sum: arr[low]}
	} else {
		mid := (high + low) / 2
		left := MaxSubArray(arr, low, mid)
		right := MaxSubArray(arr, mid+1, high)
		cross := MaxCrossSubArray(arr, low, mid, high)
		if left.sum >= right.sum && left.sum >= cross.sum {
			return left
		} else if right.sum >= left.sum && right.sum >= cross.sum {
			return right
		} else {
			return cross
		}
	}
}

// Brute Force Approach
func BFMaxSubArray(arr []int) resultTuple {
	res := resultTuple{start: 0, end: 0, sum: constants.MinInt}
	for i := 0; i < len(arr)-1; i++ {
		sum := arr[i]
		start := false
		for j := i + 1; j < len(arr); j++ {
			sum += arr[j]
			if sum > res.sum {
				if !start {
					start = true
					res.start = i
				}
				res.sum = sum
				res.end = j
			}
		}
	}
	return res
}

// Linear approach
func LMaxSubArray(arr []int) resultTuple {
	var res resultTuple = resultTuple{start: 0, end: 0, sum: constants.MinInt}
	sum := arr[0]
	for i := 1; i < len(arr)-1; i++ {
		sum += arr[i]
		if sum > res.sum {
			res.sum = sum
			res.end = i
		}
	}
	sum = arr[res.end+1]
	for i := res.end; i >= 0; i-- {
		sum += arr[i]
		if sum > res.sum {
			res.sum = sum
			res.start = i
		}
	}
	return res
}

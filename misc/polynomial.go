package misc

// arr is an array of coefficient of polynomial,
// and x is the value at which it should be evaluated

func HornerRule(arr []int, x int) int {
	y := 0
	for i := 0; i < len(arr); i++ {
		y = arr[i] + x*y
	}
	return y
}

package linear

import (
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {
	var testCases = []struct {
		name          string
		arr           []int
		want          []int
		intergerRange []int
	}{
		{
			name:          "Array containing numbers from range(0-1)",
			arr:           []int{0, 0, 1, 1, 0, 0, 1, 0},
			want:          []int{0, 0, 0, 0, 0, 1, 1, 1},
			intergerRange: []int{0, 1},
		},
	}
	for _, testCase := range testCases {
		got := CountingSort(testCase.arr, testCase.intergerRange)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("%s:\n got: %v,\n want: %v", testCase.name, got, testCase.want)
		}
	}
}

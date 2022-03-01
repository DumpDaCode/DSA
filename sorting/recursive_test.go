package sorting

import "testing"

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []int
		p   int
		r   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Sorted Array",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				p:   0,
				r:   4,
			},
		},
		{
			name: "Reverse Sorted Array",
			args: args{
				arr: []int{5, 4, 3, 2, 1},
				p:   0,
				r:   4,
			},
		},
		{
			name: "Unsorted Array",
			args: args{
				arr: []int{1, 3, 6, 2, 5},
				p:   0,
				r:   4,
			},
		},
		{
			name: "Long Array",
			args: args{
				arr: []int{1, 3, 6, 2, 5, 6, 7, 23, 1, 2, 4, 5, 6, 8, 9, 0, 2, 3, 4, 5, 6, 6, 3, 5, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.arr, tt.args.p, tt.args.r)
			t.Log(tt.args.arr)
		})
	}
}

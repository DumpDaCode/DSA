package sorting

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	type args struct {
		arr []int
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
			},
		},
		{
			name: "Reverse Sorted Array",
			args: args{
				arr: []int{5, 4, 3, 2, 1},
			},
		},
		{
			name: "Unsorted Array",
			args: args{
				arr: []int{1, 3, 6, 2, 5},
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
			InsertionSort(tt.args.arr)
			t.Log(tt.args.arr)
		})
	}
}

func TestBubbleSort(t *testing.T) {
	type args struct {
		arr []int
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
			},
		},
		{
			name: "Reverse Sorted Array",
			args: args{
				arr: []int{5, 4, 3, 2, 1},
			},
		},
		{
			name: "Unsorted Array",
			args: args{
				arr: []int{1, 3, 6, 2, 5},
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
			BubbleSort(tt.args.arr)
			t.Log(tt.args.arr)
		})
	}
}

func TestSelectionSort(t *testing.T) {
	type args struct {
		arr []int
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
			},
		},
		{
			name: "Reverse Sorted Array",
			args: args{
				arr: []int{5, 4, 3, 2, 1},
			},
		},
		{
			name: "Unsorted Array",
			args: args{
				arr: []int{1, 3, 6, 2, 5},
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
			SelectionSort(tt.args.arr)
			t.Log(tt.args.arr)
		})
	}
}

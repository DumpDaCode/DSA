package algorithms

import (
	"reflect"
	"testing"
)

func TestMaxCrossSubArray(t *testing.T) {
	type args struct {
		arr  []int
		low  int
		mid  int
		high int
	}
	tests := []struct {
		name string
		args args
		want resultTuple
	}{
		// TODO: Add test cases.
		{
			name: "Small Sub-Array",
			args: args{
				arr:  []int{1, -2, 4, -1},
				low:  0,
				mid:  1,
				high: 3,
			},
			want: resultTuple{start: 0, end: 2, sum: 3},
		},
		{
			name: "Big Sub-Array",
			args: args{
				arr:  []int{1, -2, 4, -1, 3, -6, 7, 9, 0, -1},
				low:  0,
				mid:  4,
				high: 9,
			},
			want: resultTuple{start: 2, end: 7, sum: 16},
		},
		{
			name: "Sample 3",
			args: args{
				arr:  []int{1, -2},
				low:  0,
				mid:  0,
				high: 1,
			},
			want: resultTuple{start: 0, end: 1, sum: -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxCrossSubArray(tt.args.arr, tt.args.low, tt.args.mid, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxCrossSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSubArray(t *testing.T) {
	type args struct {
		arr  []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want resultTuple
	}{
		// TODO: Add test cases.
		{
			name: "Small Sub-Array",
			args: args{
				arr:  []int{1, -2, 4, -1},
				low:  0,
				high: 3,
			},
			want: resultTuple{start: 2, end: 2, sum: 4},
		},
		{
			name: "Big Sub-Array",
			args: args{
				arr:  []int{1, -2, 4, -1, 3, -6, 7, 9, 0, -1},
				low:  0,
				high: 9,
			},
			want: resultTuple{start: 6, end: 7, sum: 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArray(tt.args.arr, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBFMaxSubArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want resultTuple
	}{
		// TODO: Add test cases.
		{
			name: "Small Sub-Array",
			args: args{
				arr: []int{1, -2, 4, -1},
			},
			want: resultTuple{start: 0, end: 2, sum: 3},
		},
		{
			name: "Big Sub-Array",
			args: args{
				arr: []int{1, -2, 4, -1, 3, -6, 7, 9, 0, -1},
			},
			want: resultTuple{start: 2, end: 7, sum: 16},
		},
		{
			name: "Sample 3",
			args: args{
				arr: []int{1, -2},
			},
			want: resultTuple{start: 0, end: 1, sum: -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BFMaxSubArray(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BFMaxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLMaxSubArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want resultTuple
	}{
		// TODO: Add test cases.
		{
			name: "Small Sub-Array",
			args: args{
				arr: []int{1, -2, 4, -1},
			},
			want: resultTuple{start: 0, end: 2, sum: 3},
		},
		{
			name: "Big Sub-Array",
			args: args{
				arr: []int{1, -2, 4, -1, 3, -6, 7, 9, 0, -1},
			},
			want: resultTuple{start: 2, end: 7, sum: 16},
		},
		{
			name: "Sample 3",
			args: args{
				arr: []int{1, -2},
			},
			want: resultTuple{start: 0, end: 1, sum: -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LMaxSubArray(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LMaxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

package heap

import (
	"testing"
)

func TestHeap_BuildMaxHeap(t *testing.T) {
	type fields struct {
		length int
		size   int
		data   []int
	}
	type args struct {
		arr []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "Simple Test",
			fields: fields{
				length: 0,
				size:   0,
				data:   []int{},
			},
			args: args{
				arr: []int{1, 2, 3, 4},
			},
		},
		{
			name: "Large array",
			fields: fields{
				length: 0,
				size:   0,
				data:   []int{},
			},
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 34},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				length: tt.fields.length,
				size:   tt.fields.size,
				data:   tt.fields.data,
			}
			h.BuildMaxHeap(tt.args.arr)
			t.Log(h.data)
		})
	}
}

func TestHeap_BuildMinHeap(t *testing.T) {
	type fields struct {
		length int
		size   int
		data   []int
	}
	type args struct {
		arr []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "Simple Test",
			fields: fields{
				length: 0,
				size:   0,
				data:   []int{},
			},
			args: args{
				arr: []int{4, 3, 2, 1},
			},
		},
		{
			name: "Large array",
			fields: fields{
				length: 0,
				size:   0,
				data:   []int{},
			},
			args: args{
				arr: []int{1, 2, 7, 8, 9, 12, 34, 3, 4, 5, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				length: tt.fields.length,
				size:   tt.fields.size,
				data:   tt.fields.data,
			}
			h.BuildMinHeap(tt.args.arr)
			t.Log(h.data)
		})
	}
}

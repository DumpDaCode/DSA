package heap

import "testing"

func TestHeap_HeapSort(t *testing.T) {
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
				arr: []int{4, 2, 3, 1},
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
			h.HeapSort(tt.args.arr)
			t.Log(h.data)
		})
	}
}

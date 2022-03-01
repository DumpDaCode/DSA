package misc

import (
	"testing"
)

func TestHornerRule(t *testing.T) {
	type args struct {
		arr []int
		x   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Simple",
			args: args{
				arr: []int{1, 4, 2, 5},
				x:   2,
			},
			want: 33,
		},
		{
			name: "With negatives",
			args: args{
				arr: []int{1, -1, 2, -3},
				x:   2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HornerRule(tt.args.arr, tt.args.x); got != tt.want {
				t.Errorf("HornerRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

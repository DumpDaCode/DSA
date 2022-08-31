package stack

import "testing"

func TestIsEmpty(t *testing.T) {
	var testCases = []struct {
		name string
		s    Stack
		want bool
	}{
		{
			name: "Stack is empty",
			s: Stack{
				Top:  0,
				Data: []int{},
			},
			want: true,
		},
		{
			name: "Stack is not empty",
			s: Stack{
				Top:  1,
				Data: []int{1},
			},
			want: false,
		},
	}

	for idx, testCase := range testCases {
		got := testCase.s.IsEmpty()
		if got != testCase.want {
			t.Errorf("Case %d(%s) returned wrong output: got(%v), want(%v)", idx, testCase.name, got, testCase.want)
		}
	}
}

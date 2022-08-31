package stack

import (
	"testing"
)

func TestStack_IsEmpty(t *testing.T) {
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

func TestStack_Push(t *testing.T) {
	var testCases = []struct {
		name    string
		s       *Stack
		wantErr bool
	}{
		{
			name:    "Pushing into stack that is not initialised",
			s:       nil,
			wantErr: true,
		},
		{
			name: "Pushing into stack that is initialised",
			s: &Stack{
				Top:  1,
				Data: []int{1},
			},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		err := testCase.s.Push(1)
		got := err != nil
		if got != testCase.wantErr {
			t.Errorf("%s: returned (%v) when user had set wantErr to (%v)", testCase.name, err, testCase.wantErr)
		}
	}
}

func TestStack_Pop(t *testing.T) {
	var testCases = []struct {
		name    string
		s       *Stack
		wantErr bool
	}{
		{
			name:    "Popping from stack that is not initialised",
			s:       nil,
			wantErr: true,
		},
		{
			name: "Popping from stack that does not contain data",
			s: &Stack{
				Top:  0,
				Data: []int{},
			},
			wantErr: true,
		},
		{
			name: "Popping from stack that is initialised",
			s: &Stack{
				Top:  1,
				Data: []int{1},
			},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		_, err := testCase.s.Pop()
		got := err != nil
		if got != testCase.wantErr {
			t.Errorf("%s: returned (%v) when user had set wantErr to (%v)", testCase.name, err, testCase.wantErr)
		}
	}
}

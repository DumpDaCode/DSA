package stack

import "errors"

type Stack struct {
	Top  int
	Data []int
}

func New() *Stack {
	s := new(Stack)
	s.Top = 0
	s.Data = make([]int, 0)
	return s
}

func (s Stack) IsEmpty() bool {
	return s.Top == 0
}

func (s Stack) Push(el int) {
	s.Top += 1
	s.Data = append(s.Data, el)
}

func (s Stack) Pop() (int, error) {
	var el int
	if s.IsEmpty() {
		return -1, errors.New("Stack is empty can't pop")
	}
	el = s.Data[s.Top]
	s.Top -= 1
	s.Data = s.Data[:s.Top]
	return el, nil
}

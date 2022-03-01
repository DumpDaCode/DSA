package stack

type AStack []string

func (a AStack) IsEmpty() bool {
	return len(a) == 0
}

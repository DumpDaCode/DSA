package queue

type Queue struct {
	Front, Back int
	Data        []string
}

func New() *Queue {
	q := new(Queue)
	q.Front, q.Back = 0, 0
	q.Data = make([]string, 0)
	return q
}

func (q *Queue) Enqueue(el int) error {
	return nil
}

func (q *Queue) Dequeue() error {
	return nil
}

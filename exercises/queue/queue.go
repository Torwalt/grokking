package queue

import "errors"

// TODO: Not yet implemented.

type Queue struct {
	data  []int
	write int
	read  int
	count int
}

func NewQueue(size int) (*Queue, error) {
	if size < 2 {
		return nil, errors.New("queue must have at least size of two")
	}

	return &Queue{
		data:  make([]int, size),
		write: 0,
		read:  0,
		count: 0,
	}, nil
}

func (q *Queue) Enqueue(e int) {
	if q.count == len(q.data) {
		q.read = 0
	}

	q.data[q.read] = e
	q.read++
	q.count++
}

func (q *Queue) Dequeue() int {
	if q.count == 0 {
		return 0
	}

	ret := q.data[q.write]
	q.count--

	return ret
}

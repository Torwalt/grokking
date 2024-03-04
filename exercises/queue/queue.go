package queue

import "errors"

type Graph map[string][]string

type Queue struct {
	data []string
}

func NewQueue(size int) (*Queue, error) {
	if size < 2 {
		return nil, errors.New("queue must have at least size of two")
	}

	return &Queue{
		data: make([]string, 0, size),
	}, nil
}

func (q *Queue) Enqueue(e ...string) {
	q.data = append(q.data, e...)
}

func (q *Queue) Dequeue() (string, bool) {
	if len(q.data) == 0 {
		return "", false
	}

	out := q.data[0]
	q.data = q.data[1:]

	return out, true
}

// The idea of BreadthFirstSearch is to find a Node in a graph in a non-greedy
// manner. When supplied a start, BFS will append adjacent Nodes of the start
// Node to a FIFO queue. Each Node in the queue will be checked whether it is
// the searched for Node and if not, the adjacent Nodes of that Node are also
// added to the queue. To prevent an infinite loop of Nodes pointing to each
// other, another structure could track the already searched Nodes.
func BreadthFirstSearch(g Graph, start, find string) bool {
	q, _ := NewQueue(10)

	if start == find {
		return true
	}

	searched := map[string]struct{}{}
	searched[start] = struct{}{}

	s, ok := g[start]
	if !ok {
		return false
	}

	q.Enqueue(s...)

	for {
		d, ok := q.Dequeue()
		if !ok {
			return false
		}

		if _, ok := searched[d]; ok {
			continue
		}

		if d == find {
			return true
		}

		searched[d] = struct{}{}

		s, ok := g[d]
		if !ok {
			continue
		}

		q.Enqueue(s...)
	}
}

package queue

import (
	"fmt"
	"sync"
)

type Q struct {
	values []int
	mu     *sync.Mutex
	n      int
}

func Queue(N int) *Q {
	return &Q{
		values: make([]int, 0, N),
		mu:     &sync.Mutex{},
		n:      N,
	}
}

func (q *Q) Push(val int) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.values) == q.n {
		return fmt.Errorf("queue full")
	}
	q.values = append(q.values, val)
	return nil
}

func (q *Q) Pop() error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.values) == 0 {
		return fmt.Errorf("empty queue")
	}
	q.values = q.values[1:]
	return nil
}

func (q *Q) Peek() (*int, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.values) == 0 {
		return nil, fmt.Errorf("empty queue")
	}
	return &q.values[0], nil
}

func (q *Q) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.values)
}

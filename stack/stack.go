package stack

import (
	"fmt"
	"sync"
)

type S struct {
	values []int
	mu     *sync.Mutex
	n      int
}

func Stack(N int) *S {
	return &S{
		values: make([]int, 0, N),
		mu:     &sync.Mutex{},
		n:      N,
	}
}

func (s *S) Push(val int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.values) == s.n {
		return fmt.Errorf("stack full")
	}
	s.values = append(s.values, val)
	return nil
}

func (s *S) Pop() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.values) == 0 {
		return fmt.Errorf("empty stack")
	}
	s.values = s.values[:len(s.values)-1]
	return nil
}

func (s *S) Peek() (*int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.values) == 0 {
		return nil, fmt.Errorf("empty stack")
	}
	return &s.values[len(s.values)-1], nil
}

func (s *S) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.values)
}

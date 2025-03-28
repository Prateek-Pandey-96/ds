package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackOperations(t *testing.T) {
	s := Stack(3) // Create a stack with capacity 3

	// Test Push
	assert.NoError(t, s.Push(1))
	assert.Equal(t, 1, s.Size())

	assert.NoError(t, s.Push(2))
	assert.NoError(t, s.Push(3))

	// Test Push when full
	assert.Error(t, s.Push(4))

	// Test Peek
	val, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 3, *val)

	// Test Pop
	assert.NoError(t, s.Pop())
	assert.Equal(t, 2, s.Size())

	// Pop remaining elements
	assert.NoError(t, s.Pop())
	assert.NoError(t, s.Pop())

	// Test Pop when empty
	assert.Error(t, s.Pop())

	// Test Peek when empty
	val, err = s.Peek()
	assert.Error(t, err)
	assert.Nil(t, val)
}

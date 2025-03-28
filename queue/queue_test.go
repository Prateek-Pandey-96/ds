package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueOperations(t *testing.T) {
	q := Queue(3) // Create a queue with capacity 3

	// Test Push
	assert.NoError(t, q.Push(1))
	assert.Equal(t, 1, q.Size())

	assert.NoError(t, q.Push(2))
	assert.NoError(t, q.Push(3))

	// Test Push when full
	assert.Error(t, q.Push(4))

	// Test Peek
	val, err := q.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 1, *val)

	// Test Pop
	assert.NoError(t, q.Pop())
	assert.Equal(t, 2, q.Size())

	// Pop remaining elements
	assert.NoError(t, q.Pop())
	assert.NoError(t, q.Pop())

	// Test Pop when empty
	assert.Error(t, q.Pop())

	// Test Peek when empty
	val, err = q.Peek()
	assert.Error(t, err)
	assert.Nil(t, val)
}

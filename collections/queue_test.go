package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueForward(t *testing.T) {
	var q Queue[int]

	assert.True(t, q.Empty())
	q.PushBack(5)
	assert.False(t, q.Empty())
	assert.Equal(t, q.Front(), 5)
	assert.Equal(t, q.Back(), 5)
	q.PushFront(7)
	assert.Equal(t, q.Front(), 7)
	assert.Equal(t, q.Back(), 5)
	q.PushBack(3)
	var results []int
	for v := range q.ForwardIter() {
		results = append(results, v)
	}
	assert.Equal(t, results, []int{7, 5, 3})
	results = nil
	for v := range q.BackIter() {
		results = append(results, v)
	}
	assert.Equal(t, results, []int{3, 5, 7})
	q.PopBack()
	assert.Equal(t, q.Back(), 5)
	q.PopBack()
	q.PopBack()
	assert.True(t, q.Empty())
	for _, v := range []int{9, 8, 7, 6} {
		q.PushFront(v)
	}
	q.PopFront()
	assert.Equal(t, q.Front(), 7)
}

package collections

import (
	"iter"
)

type node[T any] struct {
	v          T
	next, prev *node[T]
}

type Queue[T any] struct {
	head, tail *node[T]
}

func (q *Queue[T]) ForwardIter() iter.Seq[T] {
	return func(yield func(T) bool) {
		curr := q.head

		for curr != nil {
			if !yield(curr.v) {
				return
			}
			curr = curr.next

		}
	}
}

func (q *Queue[T]) BackIter() iter.Seq[T] {
	return func(yield func(T) bool) {
		curr := q.tail

		for curr != nil {
			if !yield(curr.v) {
				return
			}
			curr = curr.prev
		}
	}
}

func (q *Queue[T]) PushFront(v T) {
	n := &node[T]{
		v: v,
	}
	if q.head == nil {
		q.head = n
		q.tail = n
		return
	}

	temp := q.head
	temp.prev = n
	n.next = temp
	q.head = n
}

func (q *Queue[T]) PopBack() {
	q.tail = q.tail.prev
	if q.tail == nil {
		q.head = nil
		return
	}
	q.tail.next = nil
}

func (q *Queue[T]) PopFront() {
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
		return
	}
	q.head.prev = nil
}

func (q *Queue[T]) PushBack(v T) {
	n := &node[T]{
		v: v,
	}
	if q.tail == nil {
		q.head = n
		q.tail = n
		return
	}

	temp := q.tail
	temp.next = n
	n.prev = temp
	q.tail = n
}

func (q *Queue[T]) Empty() bool {
	return q.head == nil
}

func (q *Queue[T]) Front() T {
	return q.head.v
}

func (q *Queue[T]) Back() T {
	return q.tail.v
}

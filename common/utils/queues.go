package util

import "errors"

type Queue[T any] []T

var (
	ErrQueueEmpty = errors.New("queue empty")
)

func (q Queue[T]) Enqueue(e T) {
	q = append(q, e)
}

func (q Queue[T]) Dequeue() (T, error) {
	if len(q) > 0 {
		v := q[0]
		q = q[1:]
		return v, nil
	}

	var e T
	return e, ErrQueueEmpty
}

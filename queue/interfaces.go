package queue

import "errors"

type Item interface{}

type Queue interface {
	Empty() bool
	Enqueue(Item) error
	Dequeue() (Item, error)
	Peek() (Item, error)
}

var (
	QueueIsFullError = errors.New(
		"queue has achieved it's max capacity")
	QueueIsEmptyError = errors.New(
		"queue is empty")
)

package queue

type Item interface{}

type Queue interface {
	Empty() bool
	Enqueue(Item) error
	Dequeue() (Item, error)
	Peek() (Item, error)
}

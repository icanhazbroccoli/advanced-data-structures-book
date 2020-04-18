package queue

type ArrayQueue struct {
	items []Item
	front int
	rear  int
	size  int
}

var _ Queue = (*ArrayQueue)(nil)

func NewArrayQueue(size int) *ArrayQueue {
	return &ArrayQueue{
		items: make([]Item, size+1),
		front: 0,
		rear:  0,
		size:  size + 1,
	}
}

func (q *ArrayQueue) Empty() bool {
	return q.front == q.rear
}

func (q *ArrayQueue) Enqueue(item Item) error {
	if q.front == (q.rear+1)%q.size {
		return QueueIsFullError
	}
	q.items[q.rear] = item
	q.rear = (q.rear + 1) % q.size
	return nil
}

func (q *ArrayQueue) Dequeue() (Item, error) {
	if q.Empty() {
		return nil, QueueIsEmptyError
	}
	item := q.items[q.front]
	q.front = (q.front + 1) % q.size
	return item, nil
}

func (q *ArrayQueue) Peek() (Item, error) {
	if q.Empty() {
		return nil, QueueIsEmptyError
	}
	return q.items[q.front], nil
}

func (q *ArrayQueue) Inspect() []Item {
	res := make([]Item, 0, len(q.items))
	for i := q.front; i != q.rear; i = (i + 1) % q.size {
		res = append(res, q.items[i])
	}
	return res
}

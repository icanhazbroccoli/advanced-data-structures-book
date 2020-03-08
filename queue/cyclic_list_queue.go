package queue

type CyclicListQueue struct {
	item Item
	next *CyclicListQueue
}

var _ Queue = (*CyclicListQueue)(nil)

func NewCyclicListQueue() *CyclicListQueue {
	entrypoint := &CyclicListQueue{}
	placeholder := &CyclicListQueue{}
	entrypoint.next = placeholder
	placeholder.next = placeholder
	return entrypoint
}

func (q *CyclicListQueue) Empty() bool {
	return q.next == q.next.next
}

func (q *CyclicListQueue) Enqueue(item Item) error {
	new := &CyclicListQueue{}
	new.item = item
	tmp := q.next
	q.next = new
	new.next = tmp.next
	tmp.next = new
	return nil
}

func (q *CyclicListQueue) Dequeue() (Item, error) {
	if q.Empty() {
		return nil, QueueIsEmptyError
	}
	tmp := q.next.next.next
	q.next.next.next = tmp.next
	if tmp == q.next {
		q.next = tmp.next
	}
	return tmp.item, nil
}

func (q *CyclicListQueue) Peek() (Item, error) {
	if q.Empty() {
		return nil, QueueIsEmptyError
	}
	return q.next.next.next.item, nil
}

func (q *CyclicListQueue) Traverse() []Item {
	res := make([]Item, 0)
	start := q.next.next
	ptr := start.next
	for ptr != start {
		res = append(res, ptr.item)
		ptr = ptr.next
	}
	return res
}

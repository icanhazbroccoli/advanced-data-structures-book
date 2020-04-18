package queue

type queueNode struct {
	item Item
	next *queueNode
}

type LinkedListQueue struct {
	insert, remove *queueNode
}

var _ Queue = (*LinkedListQueue)(nil)

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{}
}

func (q *LinkedListQueue) Empty() bool {
	return q.insert == nil
}

func (q *LinkedListQueue) Enqueue(item Item) error {
	tmp := &queueNode{
		item: item,
	}
	if q.insert != nil {
		q.insert.next = tmp
		q.insert = tmp
	} else {
		q.remove = tmp
		q.insert = tmp
	}
	return nil
}

func (q *LinkedListQueue) Dequeue() (Item, error) {
	if q.Empty() {
		return nil, QueueIsEmptyError
	}
	tmp := q.remove
	q.remove = tmp.next
	if q.remove == nil {
		q.insert = nil
	}
	return tmp.item, nil
}

func (q *LinkedListQueue) Peek() (Item, error) {
	if q.Empty() {
		return nil, QueueIsEmptyError
	}
	return q.remove.item, nil
}

func (q *LinkedListQueue) Inspect() []Item {
	res := make([]Item, 0)
	tmp := q.remove
	for tmp != nil {
		res = append(res, tmp.item)
		tmp = tmp.next
	}
	return res
}

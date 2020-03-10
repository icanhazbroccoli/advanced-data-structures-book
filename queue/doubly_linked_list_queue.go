package queue

type DoublyLinkedListQueue struct {
	item       Item
	next, prev *DoublyLinkedListQueue
}

var _ Queue = (*DoublyLinkedListQueue)(nil)

func NewDoublyLinkedListQueue() *DoublyLinkedListQueue {
	q := &DoublyLinkedListQueue{}
	q.next = q
	q.prev = q
	return q
}

func (q *DoublyLinkedListQueue) Empty() bool {
	return q.next == q
}

func (q *DoublyLinkedListQueue) Enqueue(item Item) error {
	new := &DoublyLinkedListQueue{
		item: item,
	}
	new.next = q.next
	q.next = new
	new.next.prev = new
	new.prev = q
	return nil
}

func (q *DoublyLinkedListQueue) Dequeue() (Item, error) {
	if q.Empty() {
		return nil, QueueIsEmptyError
	}
	tmp := q.prev
	item := tmp.item
	tmp.prev.next = q
	q.prev = tmp.prev
	return item, nil
}

func (q *DoublyLinkedListQueue) Peek() (Item, error) {
	if q.Empty() {
		return nil, QueueIsEmptyError
	}
	return q.prev.item, nil
}

func (q *DoublyLinkedListQueue) Traverse() []Item {
	res := make([]Item, 0)
	ptr := q.prev
	for ptr != q {
		res = append(res, ptr.item)
		ptr = ptr.prev
	}
	return res
}

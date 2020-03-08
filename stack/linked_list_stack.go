package stack

type LinkedListStack struct {
	item Item
	next *LinkedListStack
}

var _ Stack = (*LinkedListStack)(nil)

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{}
}

func (s *LinkedListStack) Empty() bool {
	return s.next == nil
}

func (s *LinkedListStack) Push(item Item) {
	tmp := &LinkedListStack{}
	tmp.item = item
	tmp.next = s.next
	s.next = tmp
}

func (s *LinkedListStack) Pop() Item {
	tmp := s.next
	s.next = tmp.next
	return tmp.item
}

func (s *LinkedListStack) Peek() Item {
	return s.next.item
}

func (s *LinkedListStack) Traverse() []Item {
	ptr := s.next
	res := make([]Item, 0)
	for ptr != nil {
		res = append(res, ptr.item)
		ptr = ptr.next
	}
	return res
}

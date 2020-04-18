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

func (s *LinkedListStack) Push(item Item) error {
	tmp := &LinkedListStack{}
	tmp.item = item
	tmp.next = s.next
	s.next = tmp
	return nil
}

func (s *LinkedListStack) Pop() (Item, error) {
	if s.Empty() {
		return nil, StackIsEmptyError
	}
	tmp := s.next
	s.next = tmp.next
	return tmp.item, nil
}

func (s *LinkedListStack) Peek() (Item, error) {
	if s.Empty() {
		return nil, StackIsEmptyError
	}
	return s.next.item, nil
}

func (s *LinkedListStack) Inspect() []Item {
	ptr := s.next
	res := make([]Item, 0)
	for ptr != nil {
		res = append(res, ptr.item)
		ptr = ptr.next
	}
	return res
}

package stack

import "log"

type DynamicStack struct {
	items     []Item
	itemsCopy []Item
	size      int
	maxSize   int
	copySize  int
}

const (
	// we keep it deliberately small to test dynamic allocation
	DynamicStackDefaultSize = 2
)

var _ Stack = (*DynamicStack)(nil)

func NewDynamicStack() *DynamicStack {
	size := DynamicStackDefaultSize
	return &DynamicStack{
		items:   make([]Item, 0, size),
		maxSize: size,
		size:    0,
	}
}

func (s *DynamicStack) Empty() bool {
	return s.size == 0
}

func (s *DynamicStack) Push(item Item) error {
	s.items = append(s.items, item)
	s.size++
	sizeExceeded := float64(s.size) > 0.75*float64(s.maxSize)
	if len(s.itemsCopy) > 0 || sizeExceeded {
		additionalCopies := 4
		if len(s.itemsCopy) == 0 {
			s.itemsCopy = make([]Item, 0, 2*s.maxSize)
		}
		for additionalCopies > 0 && s.copySize < s.size {
			s.itemsCopy = append(s.itemsCopy, s.items[s.copySize])
			s.copySize++
			additionalCopies--
		}
		log.Printf("copy state: %#v", s.itemsCopy[:s.copySize])
		if s.copySize == s.size {
			s.items = s.itemsCopy
			s.maxSize *= 2
			s.itemsCopy = nil
			s.copySize = 0
			log.Printf("copy is complete: %#v", s.items[:s.size])
		}
	}
	return nil
}

func (s *DynamicStack) Pop() (Item, error) {
	if s.Empty() {
		return nil, StackIsEmptyError
	}
	s.size--
	item := s.items[s.size]
	if s.copySize == s.size {
		s.items = s.itemsCopy
		s.maxSize *= 2
		s.itemsCopy = nil
		s.copySize = 0
		log.Printf("copy is complete: %#v", s.items[:s.size])
	}
	return item, nil
}

func (s *DynamicStack) Peek() (Item, error) {
	if s.Empty() {
		return nil, StackIsEmptyError
	}
	return s.items[s.size-1], nil
}

func (s *DynamicStack) Inspect() []Item {
	return s.items[:s.size]
}

package main

import (
	"log"
	"reflect"
)

type Item int64

type Stack struct {
	item Item
	next *Stack
}

func New() *Stack {
	s := &Stack{}
	return s
}

func (s *Stack) Empty() bool {
	return s.next == nil
}

func (s *Stack) Push(item Item) {
	tmp := &Stack{}
	tmp.item = item
	tmp.next = s.next
	s.next = tmp
}

func (s *Stack) Pop() Item {
	tmp := s.next
	s.next = tmp.next
	return tmp.item
}

func (s *Stack) Peek() Item {
	return s.next.item
}

func (s *Stack) Traverse() []Item {
	ptr := s.next
	res := make([]Item, 0)
	for ptr != nil {
		res = append(res, ptr.item)
		ptr = ptr.next
	}
	return res
}

type cmdkind string

const (
	cmdempty cmdkind = "empty"
	cmdpush          = "push"
	cmdpop           = "pop"
	cmdpeek          = "peek"
)

type cmd struct {
	kind cmdkind
	args []Item
}

func main() {
	tests := []struct {
		cmd  cmd
		expv interface{}
		expt []Item
	}{
		{
			cmd:  cmd{kind: cmdempty},
			expv: true,
			expt: []Item{},
		},
		{
			cmd:  cmd{kind: cmdpush, args: []Item{42}},
			expt: []Item{42},
		},
		{
			cmd:  cmd{kind: cmdempty},
			expv: false,
			expt: []Item{42},
		},
		{
			cmd:  cmd{kind: cmdpush, args: []Item{1}},
			expt: []Item{1, 42},
		},
		{
			cmd:  cmd{kind: cmdpeek},
			expv: Item(1),
			expt: []Item{1, 42},
		},
		{
			cmd:  cmd{kind: cmdpop},
			expv: Item(1),
			expt: []Item{42},
		},
	}

	s := New()

	var v interface{}
	for _, tt := range tests {
		v = nil
		switch tt.cmd.kind {
		case cmdempty:
			v = s.Empty()
		case cmdpush:
			s.Push(tt.cmd.args[0])
		case cmdpop:
			v = s.Pop()
		case cmdpeek:
			v = s.Peek()
		}
		if v != nil {
			if !reflect.DeepEqual(v, tt.expv) {
				log.Fatalf("unexpected value returned by operation %q: got: %#v, want: %#v", tt.cmd.kind, v, tt.expv)
			}
		}
		t := s.Traverse()
		if !reflect.DeepEqual(t, tt.expt) {
			log.Fatalf("unexpected traversal: got: %#v, want: %#v", t, tt.expt)
		}
	}
}

package main

import (
	"log"
	"reflect"
)

type Item int64

const (
	blocksize = 4
)

type Stack struct {
	items [blocksize]Item
	top   int
	prev  *Stack
}

func New() *Stack {
	s := &Stack{
		top: -1,
	}
	return s
}

func (s *Stack) Empty() bool {
	return s.top == -1 && s.prev == nil
}

func (s *Stack) Push(item Item) {
	if s.top >= len(s.items)-1 {
		prev := New()
		prev.items = s.items
		prev.top = s.top
		prev.prev = s.prev
		var newitems [blocksize]Item
		s.items = newitems
		s.prev = prev
		s.top = -1
	}
	s.top++
	s.items[s.top] = item
}

func (s *Stack) Pop() Item {
	item := s.items[s.top]
	s.top--
	if s.top == -1 && s.prev != nil {
		s.items = s.prev.items
		s.top = s.prev.top
		s.prev = s.prev.prev
	}
	return item
}

func (s *Stack) Peek() Item {
	return s.items[s.top]
}

func (s *Stack) Traverse() [][]Item {
	res := make([][]Item, 0, 1)
	ptr := s
	for ptr != nil {
		if ptr.top >= 0 {
			res = append(res, ptr.items[0:ptr.top+1])
		}
		ptr = ptr.prev
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
		expt [][]Item
	}{
		{
			cmd:  cmd{kind: cmdempty},
			expv: true,
			expt: [][]Item{},
		},
		{
			cmd:  cmd{kind: cmdpush, args: []Item{42}},
			expt: [][]Item{{42}},
		},
		{
			cmd:  cmd{kind: cmdempty},
			expv: false,
			expt: [][]Item{{42}},
		},
		{
			cmd:  cmd{kind: cmdpush, args: []Item{1}},
			expt: [][]Item{{42, 1}},
		},
		{
			cmd:  cmd{kind: cmdpush, args: []Item{2}},
			expt: [][]Item{{42, 1, 2}},
		},
		{
			cmd:  cmd{kind: cmdpush, args: []Item{3}},
			expt: [][]Item{{42, 1, 2, 3}},
		},
		{
			cmd:  cmd{kind: cmdpush, args: []Item{4}},
			expt: [][]Item{{4}, {42, 1, 2, 3}},
		},
		{
			cmd:  cmd{kind: cmdpush, args: []Item{5}},
			expt: [][]Item{{4, 5}, {42, 1, 2, 3}},
		},
		{
			cmd:  cmd{kind: cmdpush, args: []Item{6}},
			expt: [][]Item{{4, 5, 6}, {42, 1, 2, 3}},
		},
		{
			cmd:  cmd{kind: cmdpeek},
			expv: Item(6),
			expt: [][]Item{{4, 5, 6}, {42, 1, 2, 3}},
		},
		{
			cmd:  cmd{kind: cmdpop},
			expv: Item(6),
			expt: [][]Item{{4, 5}, {42, 1, 2, 3}},
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

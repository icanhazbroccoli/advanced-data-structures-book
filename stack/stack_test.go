package stack

import (
	"log"
	"reflect"
	"testing"
)

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

func TestLinkedListStack(t *testing.T) {
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

	s := NewLinkedListStack()

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
				t.Errorf("unexpected value returned by operation %q: got: %#v, want: %#v", tt.cmd.kind, v, tt.expv)
			}
		}
		if tr := s.Traverse(); !reflect.DeepEqual(tr, tt.expt) {
			t.Errorf("unexpected traversal: got: %#v, want: %#v", tr, tt.expt)
		}
	}
}

func TestLinkedBlockStack(t *testing.T) {
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

	s := NewLinkedBlockStack()

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

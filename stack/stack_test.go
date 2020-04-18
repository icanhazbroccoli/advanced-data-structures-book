package stack

import (
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
	var err error
	for _, tt := range tests {
		v = nil
		err = nil
		switch tt.cmd.kind {
		case cmdempty:
			v = s.Empty()
		case cmdpush:
			err = s.Push(tt.cmd.args[0])
		case cmdpop:
			v, err = s.Pop()
		case cmdpeek:
			v, err = s.Peek()
		}
		if err != nil {
			t.Errorf("unexpected error on %q: %s", tt.cmd.kind, err)
		}
		if v != nil {
			if !reflect.DeepEqual(v, tt.expv) {
				t.Errorf("unexpected value returned by operation %q: got: %#v, want: %#v", tt.cmd.kind, v, tt.expv)
			}
		}
		if tr := s.Inspect(); !reflect.DeepEqual(tr, tt.expt) {
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
	var err error
	for _, tt := range tests {
		v = nil
		err = nil
		switch tt.cmd.kind {
		case cmdempty:
			v = s.Empty()
		case cmdpush:
			err = s.Push(tt.cmd.args[0])
		case cmdpop:
			v, err = s.Pop()
		case cmdpeek:
			v, err = s.Peek()
		}
		if err != nil {
			t.Errorf("unexpected error on %q: %s", tt.cmd.kind, err)
		}
		if v != nil {
			if !reflect.DeepEqual(v, tt.expv) {
				t.Errorf("unexpected value returned by operation %q: got: %#v, want: %#v", tt.cmd.kind, v, tt.expv)
			}
		}
		tr := s.Inspect()
		if !reflect.DeepEqual(tr, tt.expt) {
			t.Errorf("unexpected traversal: got: %#v, want: %#v", tr, tt.expt)
		}
	}
}

func TestDynamicStack(t *testing.T) {
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
			expt: []Item{42, 1},
		},
		{
			cmd:  cmd{kind: cmdpush, args: []Item{2}},
			expt: []Item{42, 1, 2},
		},
		{
			cmd:  cmd{kind: cmdpush, args: []Item{3}},
			expt: []Item{42, 1, 2, 3},
		},
		{
			cmd:  cmd{kind: cmdpush, args: []Item{4}},
			expt: []Item{42, 1, 2, 3, 4},
		},
		{
			cmd:  cmd{kind: cmdpush, args: []Item{5}},
			expt: []Item{42, 1, 2, 3, 4, 5},
		},
		{
			cmd:  cmd{kind: cmdpush, args: []Item{6}},
			expt: []Item{42, 1, 2, 3, 4, 5, 6},
		},
		{
			cmd:  cmd{kind: cmdpeek},
			expv: Item(6),
			expt: []Item{42, 1, 2, 3, 4, 5, 6},
		},
		{
			cmd:  cmd{kind: cmdpop},
			expv: Item(6),
			expt: []Item{42, 1, 2, 3, 4, 5},
		},
	}

	s := NewDynamicStack()

	var v interface{}
	var err error
	for _, tt := range tests {
		v = nil
		err = nil
		switch tt.cmd.kind {
		case cmdempty:
			v = s.Empty()
		case cmdpush:
			err = s.Push(tt.cmd.args[0])
		case cmdpop:
			v, err = s.Pop()
		case cmdpeek:
			v, err = s.Peek()
		}
		if err != nil {
			t.Errorf("unexpected error on %q: %s", tt.cmd.kind, err)
		}
		if v != nil {
			if !reflect.DeepEqual(v, tt.expv) {
				t.Errorf("unexpected value returned by operation %q: got: %#v, want: %#v", tt.cmd.kind, v, tt.expv)
			}
		}
		tr := s.Inspect()
		if !reflect.DeepEqual(tr, tt.expt) {
			t.Errorf("unexpected traversal: got: %#v, want: %#v", tr, tt.expt)
		}
	}
}

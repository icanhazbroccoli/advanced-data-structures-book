package queue

import (
	"reflect"
	"testing"
)

type cmdkind string

const (
	cmdempty   cmdkind = "empty"
	cmdenqueue         = "enqueue"
	cmddequeue         = "dequeue"
	cmdpeek            = "peek"
)

type cmd struct {
	kind cmdkind
	args []Item
}

func TestArrayQueue(t *testing.T) {
	tests := []struct {
		cmd    cmd
		expv   interface{}
		experr error
		expt   []Item
	}{
		{
			cmd:  cmd{kind: cmdempty},
			expv: true,
			expt: []Item{},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{1}},
			expt: []Item{1},
		},
		{
			cmd:  cmd{kind: cmdempty},
			expv: false,
			expt: []Item{1},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(1),
			expt: []Item{},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{1}},
			expt: []Item{1},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{2}},
			expt: []Item{1, 2},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{3}},
			expt: []Item{1, 2, 3},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{4}},
			expt: []Item{1, 2, 3, 4},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{5}},
			expt: []Item{1, 2, 3, 4, 5},
		},
		{
			cmd:    cmd{kind: cmdenqueue, args: []Item{6}},
			experr: QueueIsFullError,
			expt:   []Item{1, 2, 3, 4, 5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(1),
			expt: []Item{2, 3, 4, 5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(2),
			expt: []Item{3, 4, 5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(3),
			expt: []Item{4, 5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(4),
			expt: []Item{5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(5),
			expt: []Item{},
		},
		{
			cmd:    cmd{kind: cmddequeue},
			experr: QueueIsEmptyError,
			expt:   []Item{},
		},
	}

	q := NewArrayQueue(5)

	var v interface{}
	var err error
	for _, tt := range tests {
		v = nil
		err = nil
		switch tt.cmd.kind {
		case cmdempty:
			v = q.Empty()
		case cmdenqueue:
			err = q.Enqueue(tt.cmd.args[0])
		case cmddequeue:
			v, err = q.Dequeue()
		case cmdpeek:
			v, err = q.Peek()
		default:
			t.Fatalf("unknown command kind: %q", tt.cmd.kind)
		}
		if !errEqual(err, tt.experr) {
			t.Errorf("unexpected error on %q: got: %q, want: %q", tt.cmd.kind, err, tt.experr)
		}
		if err == nil {
			if v != nil {
				if !reflect.DeepEqual(v, tt.expv) {
					t.Errorf("unexpected value returned by operation %q: got: %#v, want: %#v", tt.cmd.kind, v, tt.expv)
				}
			}
			if tr := q.Inspect(); !reflect.DeepEqual(tr, tt.expt) {
				t.Errorf("unexpected traversal: got: %#v, want: %#v", tr, tt.expt)
			}
		}
	}
}

func TestLinkedListQueue(t *testing.T) {
	tests := []struct {
		cmd    cmd
		expv   interface{}
		experr error
		expt   []Item
	}{
		{
			cmd:  cmd{kind: cmdempty},
			expv: true,
			expt: []Item{},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{1}},
			expt: []Item{1},
		},
		{
			cmd:  cmd{kind: cmdempty},
			expv: false,
			expt: []Item{1},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(1),
			expt: []Item{},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{1}},
			expt: []Item{1},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{2}},
			expt: []Item{1, 2},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{3}},
			expt: []Item{1, 2, 3},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{4}},
			expt: []Item{1, 2, 3, 4},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{5}},
			expt: []Item{1, 2, 3, 4, 5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(1),
			expt: []Item{2, 3, 4, 5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(2),
			expt: []Item{3, 4, 5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(3),
			expt: []Item{4, 5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(4),
			expt: []Item{5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(5),
			expt: []Item{},
		},
		{
			cmd:    cmd{kind: cmddequeue},
			experr: QueueIsEmptyError,
			expt:   []Item{},
		},
	}

	q := NewLinkedListQueue()

	var v interface{}
	var err error
	for _, tt := range tests {
		v = nil
		err = nil
		switch tt.cmd.kind {
		case cmdempty:
			v = q.Empty()
		case cmdenqueue:
			err = q.Enqueue(tt.cmd.args[0])
		case cmddequeue:
			v, err = q.Dequeue()
		case cmdpeek:
			v, err = q.Peek()
		default:
			t.Fatalf("unknown command kind: %q", tt.cmd.kind)
		}
		if !errEqual(err, tt.experr) {
			t.Errorf("unexpected error on %q: got: %q, want: %q", tt.cmd.kind, err, tt.experr)
		}
		if err == nil {
			if v != nil {
				if !reflect.DeepEqual(v, tt.expv) {
					t.Errorf("unexpected value returned by operation %q: got: %#v, want: %#v", tt.cmd.kind, v, tt.expv)
				}
			}
			if tr := q.Inspect(); !reflect.DeepEqual(tr, tt.expt) {
				t.Errorf("unexpected traversal: got: %#v, want: %#v", tr, tt.expt)
			}
		}
	}
}

func TestCyclicListQueue(t *testing.T) {
	tests := []struct {
		cmd    cmd
		expv   interface{}
		experr error
		expt   []Item
	}{
		{
			cmd:  cmd{kind: cmdempty},
			expv: true,
			expt: []Item{},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{1}},
			expt: []Item{1},
		},
		{
			cmd:  cmd{kind: cmdempty},
			expv: false,
			expt: []Item{1},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(1),
			expt: []Item{},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{1}},
			expt: []Item{1},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{2}},
			expt: []Item{1, 2},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{3}},
			expt: []Item{1, 2, 3},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{4}},
			expt: []Item{1, 2, 3, 4},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{5}},
			expt: []Item{1, 2, 3, 4, 5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(1),
			expt: []Item{2, 3, 4, 5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(2),
			expt: []Item{3, 4, 5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(3),
			expt: []Item{4, 5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(4),
			expt: []Item{5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(5),
			expt: []Item{},
		},
		{
			cmd:    cmd{kind: cmddequeue},
			experr: QueueIsEmptyError,
			expt:   []Item{},
		},
	}

	q := NewCyclicListQueue()

	var v interface{}
	var err error
	for _, tt := range tests {
		v = nil
		err = nil
		switch tt.cmd.kind {
		case cmdempty:
			v = q.Empty()
		case cmdenqueue:
			err = q.Enqueue(tt.cmd.args[0])
		case cmddequeue:
			v, err = q.Dequeue()
		case cmdpeek:
			v, err = q.Peek()
		default:
			t.Fatalf("unknown command kind: %q", tt.cmd.kind)
		}
		if !errEqual(err, tt.experr) {
			t.Errorf("unexpected error on %q: got: %q, want: %q", tt.cmd.kind, err, tt.experr)
		}
		if err == nil {
			if v != nil {
				if !reflect.DeepEqual(v, tt.expv) {
					t.Errorf("unexpected value returned by operation %q: got: %#v, want: %#v", tt.cmd.kind, v, tt.expv)
				}
			}
			if tr := q.Inspect(); !reflect.DeepEqual(tr, tt.expt) {
				t.Errorf("unexpected traversal: got: %#v, want: %#v", tr, tt.expt)
			}
		}
	}
}

func TestDoublyLinkedListQueue(t *testing.T) {
	tests := []struct {
		cmd    cmd
		expv   interface{}
		experr error
		expt   []Item
	}{
		{
			cmd:  cmd{kind: cmdempty},
			expv: true,
			expt: []Item{},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{1}},
			expt: []Item{1},
		},
		{
			cmd:  cmd{kind: cmdempty},
			expv: false,
			expt: []Item{1},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(1),
			expt: []Item{},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{1}},
			expt: []Item{1},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{2}},
			expt: []Item{1, 2},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{3}},
			expt: []Item{1, 2, 3},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{4}},
			expt: []Item{1, 2, 3, 4},
		},
		{
			cmd:  cmd{kind: cmdenqueue, args: []Item{5}},
			expt: []Item{1, 2, 3, 4, 5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(1),
			expt: []Item{2, 3, 4, 5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(2),
			expt: []Item{3, 4, 5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(3),
			expt: []Item{4, 5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(4),
			expt: []Item{5},
		},
		{
			cmd:  cmd{kind: cmddequeue},
			expv: Item(5),
			expt: []Item{},
		},
		{
			cmd:    cmd{kind: cmddequeue},
			experr: QueueIsEmptyError,
			expt:   []Item{},
		},
	}

	q := NewDoublyLinkedListQueue()

	var v interface{}
	var err error
	for _, tt := range tests {
		v = nil
		err = nil
		switch tt.cmd.kind {
		case cmdempty:
			v = q.Empty()
		case cmdenqueue:
			err = q.Enqueue(tt.cmd.args[0])
		case cmddequeue:
			v, err = q.Dequeue()
		case cmdpeek:
			v, err = q.Peek()
		default:
			t.Fatalf("unknown command kind: %q", tt.cmd.kind)
		}
		if !errEqual(err, tt.experr) {
			t.Errorf("unexpected error on %q: got: %q, want: %q", tt.cmd.kind, err, tt.experr)
		}
		if err == nil {
			if v != nil {
				if !reflect.DeepEqual(v, tt.expv) {
					t.Errorf("unexpected value returned by operation %q: got: %#v, want: %#v", tt.cmd.kind, v, tt.expv)
				}
			}
			if tr := q.Inspect(); !reflect.DeepEqual(tr, tt.expt) {
				t.Errorf("unexpected traversal: got: %#v, want: %#v", tr, tt.expt)
			}
		}
	}
}

func errEqual(e1, e2 error) bool {
	if e1 == nil || e2 == nil {
		return e1 == e2
	}
	return e1.Error() == e2.Error()
}

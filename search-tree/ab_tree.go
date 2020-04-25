package search_tree

import (
	"bytes"
	"fmt"
	"strings"
)

type ABTree struct {
	A, B   int
	keys   []SearchKey
	values []StoredObject
	next   []*ABTree
	degree int
	height int
}

var _ SearchTree = (*ABTree)(nil)

func NewABTree(A, B int) *ABTree {
	return &ABTree{
		A:      A,
		B:      B,
		keys:   make([]SearchKey, 0, B),
		values: make([]StoredObject, 0, B),
		next:   make([]*ABTree, 0, B),
		degree: 0,
		height: 0,
	}
}

func (t *ABTree) Find(key SearchKey) (StoredObject, FindStatus) {
	cur := t
	for cur.height >= 0 {
		lower := 0
		upper := cur.degree
		for upper > lower+1 {
			if key.LessThan(cur.keys[(upper+lower)/2]) {
				upper = (upper + lower) / 2
			} else {
				lower = (upper + lower) / 2
			}
		}
		if cur.height > 0 {
			cur = cur.next[lower]
		} else {
			if cur.keys[lower].EqualsTo(key) {
				return cur.values[lower], FindOk
			}
			break
		}
	}
	return nil, FindNone
}

// resize is a helper routine to adjust the slice sizes
// Overall, there is nothing terrible in keeping things in static arrays as all
// operatios use degree value to establish an active area of these slices. I'm
// using it in order to clean up the garbage and make the implementation more
// testable.
func (t *ABTree) resize(size int) {
	t.keys = t.keys[:size]
	t.next = t.next[:size]
	t.values = t.values[:size]
}

func (t *ABTree) Insert(key SearchKey, value StoredObject) InsertStatus {
	if t.isEmpty() {
		t.resize(1)
		t.keys[0] = key
		t.values[0] = value
		t.degree = 1
		return InsertOk
	}
	cur := t
	stack := []*ABTree{}
	for !cur.isLeaf() {
		stack = append(stack, cur)
		lower, upper := 0, cur.degree
		for upper > lower+1 {
			if key.LessThan(cur.keys[(upper+lower)/2]) {
				upper = (upper + lower) / 2
			} else {
				lower = (upper + lower) / 2
			}
			cur = cur.next[lower]
		}
	}
	finished := false
	var insertKey = key
	var insertNode *ABTree
	for !finished {
		start, i := 0, 0
		if !cur.isLeaf() {
			start = 1
		}
		if cur.degree < cur.B {
			i = cur.degree
			cur.resize(cur.degree + 1)
			for (i > start) && key.LessThan(cur.keys[i-1]) {
				cur.keys[i] = cur.keys[i-1]
				cur.next[i] = cur.next[i-1]
				cur.values[i] = cur.values[i-1]
				i -= 1
			}
			cur.keys[i] = insertKey
			if cur.isLeaf() {
				cur.next[i] = nil
				cur.values[i] = value
			} else {
				cur.next[i] = insertNode
				cur.values[i] = nil
			}
			cur.degree++
			finished = true
		} else {
			// the node is full, splitting
			newNode := NewABTree(cur.A, cur.B)
			var j int
			var done bool
			i, j = cur.B-1, (cur.B-1)/2
			newNode.resize(j + 1)
			for j >= 0 { // handling the upper half
				if done || (key.LessThan(cur.keys[i])) {
					newNode.keys[j] = cur.keys[i]
					newNode.next[j] = cur.next[i]
					newNode.values[j] = cur.values[i]
					i--
					j--
				} else {
					newNode.keys[j] = key
					newNode.next[j] = nil
					newNode.values[j] = value
					j--
					done = true
				}
			}
			for !done { // handling the lower half if necessary
				if key.LessThan(cur.keys[i]) && i >= start {
					cur.next[i+1] = cur.next[i]
					cur.keys[i+1] = cur.keys[i]
					cur.values[i+1] = cur.values[i]
					i--
				} else {
					cur.values[i+1] = value
					cur.keys[i+1] = key
					cur.next[i+1] = nil
					done = true
				}
			} // insertion is complete now
			cur.degree = cur.B + 1 - (cur.B+1)/2
			cur.resize(cur.degree)
			newNode.degree = (cur.B + 1) / 2
			newNode.height = cur.height
			insertNode = newNode
			insertKey = newNode.keys[0]
			// split is complete, inserting the new node above
			if len(stack) > 0 {
				// move 1 level up
				stack, cur = stack[:len(stack)-1], stack[len(stack)-1]
			} else {
				newNode = NewABTree(cur.A, cur.B)
				newNode.resize(cur.degree)
				for i := 0; i < cur.degree; i++ {
					newNode.next[i] = cur.next[i]
					newNode.keys[i] = cur.keys[i]
					newNode.values[i] = cur.values[i]
				}
				newNode.height = cur.height
				newNode.degree = cur.degree
				cur.height++
				cur.degree = 2
				cur.resize(cur.degree)
				cur.keys[0] = nil
				cur.next[0] = newNode
				cur.values[0] = nil // ???
				cur.keys[1] = insertKey
				cur.next[1] = insertNode
				cur.values[1] = nil // ???
				finished = true
			} // root split is complete
		} // node split is complete
	}
	return InsertOk
}

func (t *ABTree) Delete(key SearchKey) (StoredObject, DeleteStatus) {
	return nil, DeleteNone
}

func (t *ABTree) String() string {
	return t.toString("")
}

func (t *ABTree) toString(padding string) string {
	var out bytes.Buffer
	if t == nil {
		out.WriteString(padding + "nil")
	} else {
		children := make([]string, 0, t.degree)
		for i := 0; i < t.degree; i++ {
			children = append(children, t.next[i].toString("  "))
		}
		out.WriteString(fmt.Sprintf(`[
%s  A: %d, B: %d, degree: %d, height: %d,
`, padding, t.A, t.B, t.degree, t.height))

		out.WriteString(fmt.Sprintf("%s  keys: %+v,\n", padding, t.keys))
		out.WriteString(fmt.Sprintf("%s  values: %+v,\n", padding, t.values))
		out.WriteString(fmt.Sprintf(`%s  next: [
%[1]s  %s
%[1]s  ],
`, padding, strings.Join(children, fmt.Sprintf(",\n%s  ", padding))))
		out.WriteString(padding + "]")
	}
	return out.String()
}

func (t *ABTree) isEmpty() bool {
	return t.isLeaf() && t.degree == 0
}

func (t *ABTree) isLeaf() bool {
	return t.height == 0
}

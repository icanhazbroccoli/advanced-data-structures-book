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
	if size > t.B {
		panic("tree fixed size exceeded on resize")
	}
	// We're keeping these dynamic resizers here in place for the sake of
	// handling non-constructor initialized trees (which is still a valid case).
	if cap(t.keys) < size {
		keys := make([]SearchKey, size, t.B)
		copy(keys, t.keys)
		t.keys = keys
	}
	t.keys = t.keys[:size]
	if cap(t.next) < size {
		next := make([]*ABTree, size, t.B)
		copy(next, t.next)
		t.next = next
	}
	t.next = t.next[:size]
	if cap(t.values) < size {
		values := make([]StoredObject, size, t.B)
		copy(values, t.values)
		t.values = values
	}
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
	for i := 0; i < cur.degree; i++ {
		if cur.keys[i].EqualsTo(key) {
			cur.values[i] = value
			return InsertNone
		} else if key.LessThan(cur.keys[i]) {
			break
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
				cur.values[0] = nil
				cur.keys[1] = insertKey
				cur.next[1] = insertNode
				cur.values[1] = nil
				finished = true
			} // root split is complete
		} // node split is complete
	}
	return InsertOk
}

func (t *ABTree) Delete(key SearchKey) (StoredObject, DeleteStatus) {
	cur := t
	stack := []*ABTree{}
	indexStack := []int{}
	for !cur.isLeaf() {
		lower, upper := 0, cur.degree
		for upper > lower+1 {
			if key.LessThan(cur.keys[(lower+upper)/2]) {
				upper = (lower + upper) / 2
			} else {
				lower = (lower + upper) / 2
			}
			stack = append(stack, cur)
			indexStack = append(indexStack, lower)
			cur = cur.next[lower]
		}
	}
	index := -1
	for i := 0; i < cur.degree; i++ {
		if key.LessThanOrEqualsTo(cur.keys[i]) {
			if cur.keys[i].EqualsTo(key) {
				index = i
			}
			break
		}
	}
	if index == -1 {
		return nil, DeleteNone
	}
	delObj := cur.values[index]
	cur.degree--
	for index < cur.degree {
		cur.keys[index] = cur.keys[index+1]
		cur.values[index] = cur.values[index+1]
		index++
	}
	cur.resize(cur.degree)

	for {
		if cur.degree >= cur.A { // the node is full enough
			break
		}
		if len(stack) == 0 {
			// we are dealing with the root
			if cur.degree >= 2 {
				break // we still need the root around
			} else if cur.height == 0 {
				break // deleting last keys from the root
			} else {
				tmp := cur.next[0]
				cur.resize(tmp.degree)
				for i := 0; i < tmp.degree; i++ {
					cur.next[i] = tmp.next[i]
					cur.keys[i] = tmp.keys[i]
					cur.values[i] = tmp.values[i]
				}
				cur.degree = tmp.degree
				cur.height = tmp.height
				cur.resize(cur.degree)
			}
			break
		}

		var upper *ABTree
		var ix int
		stack, upper = stack[:len(stack)-1], stack[len(stack)-1]
		indexStack, ix = indexStack[:len(indexStack)-1], indexStack[len(indexStack)-1]
		if ix < upper.degree-1 {
			neighbor := upper.next[ix+1]
			if neighbor.degree > cur.A { // sharing is possible
				index = cur.degree
				if cur.height > 0 {
					cur.keys[index] = upper.keys[ix+1]
				} else {
					cur.keys[index] = neighbor.keys[0]
					neighbor.keys[0] = neighbor.keys[1]
				}
				cur.next[index] = neighbor.next[0]
				cur.values[index] = neighbor.values[0]
				upper.keys[ix+1] = neighbor.keys[1]
				upper.values[ix+1] = neighbor.values[1]
				neighbor.next[0] = neighbor.next[1]
				for j := 2; j < neighbor.degree; j++ {
					neighbor.next[j-1] = neighbor.next[j]
					neighbor.keys[j-1] = neighbor.keys[j]
					neighbor.values[j-1] = neighbor.values[j]
				}
				neighbor.degree--
				neighbor.resize(neighbor.degree)
				cur.degree++
				cur.resize(cur.degree)
				break
				// sharing is complete
			} else {
				// must join
				index = cur.degree
				if cur.height > 0 {
					cur.keys[index] = upper.keys[ix+1]
				} else {
					// leaf level
					cur.keys[index] = neighbor.keys[0]
				}
				cur.next[index] = neighbor.next[0]
				cur.values[index] = neighbor.values[0]
				for j := 1; j < neighbor.degree; j++ {
					index++
					cur.next[index] = neighbor.next[j]
					cur.keys[index] = neighbor.keys[j]
					cur.values[index] = neighbor.values[j]
				}
				cur.degree = index + 1
				cur.resize(cur.degree)
				upper.degree--
				upper.resize(upper.degree)
				index = ix + 1
				for index < upper.degree {
					upper.next[index] = upper.next[index+1]
					upper.keys[index] = upper.keys[index+1]
					upper.values[index] = upper.values[index+1]
					index++
				}
				cur = upper
			}
		} else {
			neighbor := upper.next[ix-1]
			if neighbor.degree > cur.A {
				// sharing is possible
				cur.resize(cur.degree + 1)
				for j := cur.degree; j > 1; j-- {
					cur.next[j] = cur.next[j-1]
					cur.keys[j] = cur.keys[j-1]
					cur.values[j] = cur.values[j-1]
				}
				cur.degree = cur.degree + 1
				cur.next[1] = cur.next[0]
				cur.values[1] = cur.values[0]
				index = neighbor.degree
				cur.next[0] = neighbor.next[index-1]
				cur.values[0] = neighbor.values[index-1]
				if cur.height > 0 {
					cur.keys[1] = upper.keys[ix]
				} else {
					cur.keys[1] = cur.keys[0]
					cur.keys[0] = neighbor.keys[index-1]
				}
				upper.keys[ix] = neighbor.keys[index-1]
				neighbor.degree--
				neighbor.resize(neighbor.degree)
				break
			} else {
				// must join
				index = neighbor.degree
				neighbor.resize(neighbor.degree + cur.degree)
				if cur.height > 0 {
					neighbor.keys[index] = upper.keys[ix]
				} else {
					neighbor.keys[index] = cur.keys[0]
				}
				neighbor.next[index] = cur.next[0]
				neighbor.values[index] = cur.values[0]
				for j := 1; j < cur.degree; j++ {
					index++
					neighbor.next[index] = cur.next[j]
					neighbor.keys[index] = cur.keys[j]
					neighbor.values[index] = cur.values[j]
				}
				neighbor.degree = index + 1
				upper.degree--
				cur = upper
				upper.resize(upper.degree)
			}
		}
	}

	return delObj, DeleteOk
}

func (t *ABTree) String() string {
	noPadding := ""
	return t.toString(noPadding)
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
`, padding, strings.Join(children, ",\n"+padding+"  ")))
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

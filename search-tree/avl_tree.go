package search_tree

import (
	"bytes"
	"fmt"
)

type AVLTree struct {
	key         SearchKey
	value       StoredObject
	left, right *AVLTree
	height      int
}

var _ SearchTree = (*AVLTree)(nil)

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func (t *AVLTree) rotateLeft() {
	if t.isLeaf() {
		return
	}
	tmpNode := t.left
	tmpKey := t.key
	t.left = t.right
	t.key = t.right.key
	t.right = t.left.right
	t.left.right = t.left.left
	t.left.left = tmpNode
	t.left.key = tmpKey
}

func (t *AVLTree) rotateRight() {
	if t.isLeaf() {
		return
	}
	tmpNode := t.right
	tmpKey := t.key
	t.right = t.left
	t.key = t.left.key
	t.left = t.right.left
	t.right.left = t.right.right
	t.right.right = tmpNode
	t.right.key = tmpKey
}

func (t *AVLTree) Find(key SearchKey) (StoredObject, FindStatus) {
	if t.isEmpty() {
		return nil, FindNone
	}
	tmp := t
	for !tmp.isLeaf() {
		if key.LessThan(tmp.key) {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
	}
	if tmp.key.EqualsTo(key) {
		return tmp.value, FindOk
	}
	return nil, FindNone
}

func (t *AVLTree) Insert(key SearchKey, value StoredObject) InsertStatus {
	if t.isEmpty() {
		t.key = key
		t.value = value
		t.height = 0
		return InsertOk
	}
	tmp := t
	stack := []*AVLTree{}
	for !tmp.isLeaf() {
		stack = append(stack, tmp)
		if key.LessThan(tmp.key) {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
	}
	if tmp.key.EqualsTo(key) {
		tmp.value = value
		return InsertNone
	}
	oldLeaf := &AVLTree{
		key:    tmp.key,
		value:  tmp.value,
		height: 0,
	}
	newLeaf := &AVLTree{
		key:    key,
		value:  value,
		height: 0,
	}
	tmp.value = nil
	if tmp.key.LessThan(key) {
		tmp.left = oldLeaf
		tmp.right = newLeaf
		tmp.key = key
	} else {
		tmp.left = newLeaf
		tmp.right = oldLeaf
	}
	tmp.height = 1

	t.rebalance(stack)

	return InsertOk
}

func (t *AVLTree) Delete(key SearchKey) (StoredObject, DeleteStatus) {
	var tmp, other, upper *AVLTree
	var value StoredObject
	if t.isEmpty() {
		return nil, DeleteNone
	}
	if t.isLeaf() {
		if t.key.EqualsTo(key) {
			value = t.value
			t.key = nil
			t.value = nil
			return value, DeleteOk
		}
		return nil, DeleteNone
	}
	tmp = t
	stack := []*AVLTree{}
	for !tmp.isLeaf() {
		stack = append(stack, tmp)
		upper = tmp
		if key.LessThan(tmp.key) {
			tmp = upper.left
			other = upper.right
		} else {
			tmp = upper.right
			other = upper.left
		}
	}
	if !tmp.key.EqualsTo(key) {
		return nil, DeleteNone
	}
	upper.key = other.key
	upper.value = other.value
	upper.left = other.left
	upper.right = other.right
	upper.height = 0
	value = tmp.value

	t.rebalance(stack)

	return value, DeleteOk
}

func (*AVLTree) rebalance(stack []*AVLTree) {
	var tmp *AVLTree
	for len(stack) > 0 {
		var tmpHeight, oldHeight int
		stack, tmp = stack[:len(stack)-1], stack[len(stack)-1]
		if tmp.isLeaf() {
			continue
		}
		oldHeight = tmp.height
		if tmp.left.height-tmp.right.height == 2 {
			if tmp.left.left.height-tmp.right.height == 1 {
				tmp.rotateRight()
				tmp.right.height = tmp.right.left.height + 1
				tmp.height = tmp.right.height + 1
			} else {
				tmp.left.rotateLeft()
				tmp.rotateRight()
				tmpHeight = tmp.left.left.height
				tmp.left.height = tmpHeight + 1
				tmp.right.height = tmpHeight + 1
				tmp.height = tmpHeight + 2
			}
		} else if tmp.left.height-tmp.right.height == -2 {
			if tmp.right.right.height-tmp.left.height == 1 {
				tmp.rotateLeft()
				tmp.left.height = tmp.left.right.height + 1
				tmp.height = tmp.left.height + 1
			} else {
				tmp.right.rotateRight()
				tmp.rotateLeft()
				tmpHeight = tmp.right.right.height
				tmp.left.height = tmpHeight + 1
				tmp.right.height = tmpHeight + 1
				tmp.height = tmpHeight + 1
			}
		} else {
			if tmp.left.height > tmp.right.height {
				tmp.height = tmp.left.height + 1
			} else {
				tmp.height = tmp.right.height + 1
			}
		}
		if tmpHeight == oldHeight {
			break
		}
	}
}

func (t *AVLTree) Inspect() string {
	var out bytes.Buffer

	if t == nil {
		out.WriteString("nil")
	} else if t.isLeaf() {
		out.WriteString("(")
		out.WriteString("leaf/")
		out.WriteString(fmt.Sprintf("k:%v", t.key))
		out.WriteString("/")
		out.WriteString(fmt.Sprintf("h:%d", t.height))
		out.WriteString("/")
		out.WriteString(fmt.Sprintf("v:%v", t.value))
		out.WriteString(")")
	} else {
		out.WriteString("[")
		out.WriteString(fmt.Sprintf("k:%v", t.key))
		out.WriteString("/")
		out.WriteString(fmt.Sprintf("h:%d", t.height))
		out.WriteString("/")
		out.WriteString(fmt.Sprintf("l:%s", t.left.Inspect()))
		out.WriteString("/")
		out.WriteString(fmt.Sprintf("r:%s", t.right.Inspect()))
		out.WriteString("]")
	}

	return out.String()
}

func (t *AVLTree) isLeaf() bool {
	return t.left == nil
}

func (t *AVLTree) isEmpty() bool {
	return t.isLeaf() && t.value == nil
}

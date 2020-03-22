package search_tree

import (
	"bytes"
	"fmt"
)

type BinaryTree struct {
	key         SearchKey
	value       StoredObject
	left, right *BinaryTree
}

var _ (SearchTree) = (*BinaryTree)(nil)

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) RotateLeft() {
	tmpNode := t.left
	tmpKey := t.key
	t.left = t.right
	t.key = t.right.key
	t.right = t.left.right
	t.left.right = t.left.left
	t.left.left = tmpNode
	t.left.key = tmpKey
}

func (t *BinaryTree) RotateRight() {
	tmpNode := t.right
	tmpKey := t.key
	t.right = t.left
	t.key = t.left.key
	t.left = t.right.left
	t.right.left = t.right.right
	t.right.right = tmpNode
	t.right.key = tmpKey
}

func (t *BinaryTree) Find(key SearchKey) (StoredObject, FoundStatus) {
	if t.left == nil {
		return nil, FoundNone
	}
	tmp := t
	for !isLeaf(tmp) {
		if key.LessThan(tmp.key) {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
	}
	if tmp.key.EqualsTo(key) {
		return tmp.value, FoundOk
	}
	return nil, FoundNone
}

func isLeaf(t *BinaryTree) bool {
	return t.left == nil
}

func (t *BinaryTree) Insert(key SearchKey, value StoredObject) InsertStatus {
	if t == nil {
		return InsertNone
	}
	if isLeaf(t) {
		t.value = value
		t.key = key
		t.right = nil
	}
	tmp := t
	for !isLeaf(tmp) {
		if key.LessThan(tmp.key) {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
	}
	if tmp.key.EqualsTo(key) {
		return InsertNone
	}
	oldLeaf := &BinaryTree{
		key:   tmp.key,
		value: tmp.value,
	}
	newLeaf := &BinaryTree{
		key:   key,
		value: value,
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
	return InsertOk
}

func (t *BinaryTree) Delete(key SearchKey) (StoredObject, DeleteStatus) {
	var tmp, other, upper *BinaryTree
	var value StoredObject
	if t == nil {
		return nil, DeleteNone
	}
	if isLeaf(t) {
		if t.key.EqualsTo(key) {
			value = t.value
			t.value = nil
			return value, DeleteOk
		}
		return nil, DeleteNone
	}
	tmp = t
	for !isLeaf(tmp) {
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
	value = tmp.value
	return value, DeleteOk
}

func (t *BinaryTree) Traverse() string {
	var out bytes.Buffer

	if isLeaf(t) {
		out.WriteString("(")
		out.WriteString("leaf/")
		out.WriteString(fmt.Sprintf("k:%v", t.key))
		out.WriteString("/")
		out.WriteString(fmt.Sprintf("v:%v", t.value))
		out.WriteString(")")
	} else {
		out.WriteString("[")
		out.WriteString(fmt.Sprintf("k:%v", t.key))
		out.WriteString("/")
		out.WriteString(fmt.Sprintf("l:%s", t.left.Traverse()))
		out.WriteString("/")
		out.WriteString(fmt.Sprintf("r:%s", t.right.Traverse()))
		out.WriteString("]")
	}

	return out.String()
}

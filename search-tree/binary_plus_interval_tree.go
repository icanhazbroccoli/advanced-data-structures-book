package search_tree

import (
	"bytes"
	"fmt"
)

type BinaryPlusIntervalTree struct {
	key         SearchKey
	value       StoredObject
	left, right *BinaryPlusIntervalTree
	next, prev  *BinaryPlusIntervalTree
}

var _ (IntervalSearchTree) = (*BinaryPlusIntervalTree)(nil)

func NewBinaryPlusIntervalTree() *BinaryPlusIntervalTree {
	return &BinaryPlusIntervalTree{}
}

func (t *BinaryPlusIntervalTree) Find(key SearchKey) (StoredObject, FindStatus) {
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

func (t *BinaryPlusIntervalTree) FindInterval(a, b SearchKey) []StoredObject {
	res := []StoredObject{}
	if t.isEmpty() {
		return res
	}
	tmp := t
	for !tmp.isLeaf() {
		if tmp.key.LessThanOrEqualsTo(a) {
			tmp = tmp.right
		} else if b.LessThanOrEqualsTo(b) {
			tmp = tmp.left
		} else {
			tmp = tmp.left
		}
	}
	for tmp != nil {
		if a.LessThanOrEqualsTo(tmp.key) && tmp.key.LessThan(b) {
			res = append(res, tmp.value)
		} else if b.LessThanOrEqualsTo(tmp.key) {
			break
		}
		tmp = tmp.next
	}
	return res
}

func (t *BinaryPlusIntervalTree) Insert(key SearchKey, value StoredObject) InsertStatus {
	if t.isEmpty() {
		t.value = value
		t.key = key
		return InsertOk
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
		tmp.value = value
		return InsertNone
	}
	oldLeaf := &BinaryPlusIntervalTree{
		key:   tmp.key,
		value: tmp.value,
	}
	newLeaf := &BinaryPlusIntervalTree{
		key:   key,
		value: value,
	}
	tmp.value = nil
	if tmp.key.LessThan(key) {
		tmp.left = oldLeaf
		tmp.right = newLeaf
		tmp.key = key
		oldLeaf.prev = tmp.prev
		oldLeaf.next = newLeaf
		newLeaf.prev = oldLeaf
		newLeaf.next = tmp.next
		if oldLeaf.prev != nil {
			oldLeaf.prev.next = oldLeaf
		}
		if newLeaf.next != nil {
			newLeaf.next.prev = newLeaf
		}
	} else {
		tmp.left = newLeaf
		tmp.right = oldLeaf
		newLeaf.prev = tmp.prev
		newLeaf.next = oldLeaf
		oldLeaf.prev = newLeaf
		oldLeaf.next = tmp.next
		if newLeaf.prev != nil {
			newLeaf.prev.next = newLeaf
		}
		if oldLeaf.next != nil {
			oldLeaf.next.prev = oldLeaf
		}
	}
	tmp.next = nil
	tmp.prev = nil

	return InsertOk
}

func (t *BinaryPlusIntervalTree) Delete(key SearchKey) (StoredObject, DeleteStatus) {
	if t.isEmpty() {
		return nil, DeleteNone
	}
	if t.isLeaf() {
		if t.key.EqualsTo(key) {
			value := t.value
			t.key = nil
			t.value = nil
			return value, DeleteOk
		}
		return nil, DeleteNone
	}
	var tmp, other, upper *BinaryPlusIntervalTree
	tmp = t
	for !tmp.isLeaf() {
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
	if tmp.prev != nil {
		tmp.prev.next = tmp.next
	}
	if tmp.next != nil {
		tmp.next.prev = tmp.prev
	}
	upper.key = other.key
	upper.value = other.value
	upper.left = other.left
	upper.right = other.right
	upper.prev = other.prev
	upper.next = other.next
	value := tmp.value
	return value, DeleteOk
}

func (t *BinaryPlusIntervalTree) Inspect() string {
	var out bytes.Buffer

	if t == nil {
		out.WriteString("nil")
	} else if t.isLeaf() {
		out.WriteString("(")
		out.WriteString("leaf/")
		out.WriteString(fmt.Sprintf("k:%v", t.key))
		out.WriteString("/")
		out.WriteString(fmt.Sprintf("v:%v", t.value))
		if t.prev != nil {
			out.WriteString("/")
			out.WriteString(fmt.Sprintf("%v<-", t.prev.key))
		}
		if t.next != nil {
			out.WriteString("/")
			out.WriteString(fmt.Sprintf("->%v", t.next.key))
		}
		out.WriteString(")")
	} else {
		out.WriteString("[")
		out.WriteString(fmt.Sprintf("k:%v", t.key))
		out.WriteString("/")
		out.WriteString(fmt.Sprintf("l:%s", t.left.Inspect()))
		out.WriteString("/")
		out.WriteString(fmt.Sprintf("r:%s", t.right.Inspect()))
		out.WriteString("]")
	}

	return out.String()
}

func (t *BinaryPlusIntervalTree) isLeaf() bool {
	return t.left == nil
}

func (t *BinaryPlusIntervalTree) isEmpty() bool {
	return t.isLeaf() && t.value == nil
}

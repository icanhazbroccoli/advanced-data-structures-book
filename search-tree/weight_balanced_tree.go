package search_tree

import (
	"bytes"
	"fmt"
)

const (
	Alpha   float64 = 0.288
	Epsilon         = 0.005
)

type WeightBalancedTree struct {
	key         SearchKey
	value       StoredObject
	left, right *WeightBalancedTree
	weight      float64
}

var _ SearchTree = (*WeightBalancedTree)(nil)

func NewWeightBalancedTree() *WeightBalancedTree {
	return &WeightBalancedTree{
		weight: 0.0,
	}
}

func (t *WeightBalancedTree) rotateLeft() {
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

func (t *WeightBalancedTree) rotateRight() {
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

func (t *WeightBalancedTree) Find(key SearchKey) (StoredObject, FindStatus) {
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

func (t *WeightBalancedTree) Insert(key SearchKey, value StoredObject) InsertStatus {
	if t.isEmpty() {
		t.key = key
		t.value = value
		t.weight = 1.0
		return InsertOk
	}
	tmp := t
	stack := []*WeightBalancedTree{}
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
	oldLeaf := &WeightBalancedTree{
		value:  tmp.value,
		key:    tmp.key,
		weight: 1.0,
	}
	newLeaf := &WeightBalancedTree{
		value:  value,
		key:    key,
		weight: 1.0,
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
	tmp.weight = 2.0

	t.rebalance(stack)

	return InsertOk
}

func (t *WeightBalancedTree) Delete(key SearchKey) (StoredObject, DeleteStatus) {
	var tmp, other, upper *WeightBalancedTree
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
	stack := []*WeightBalancedTree{}
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
	upper.weight = 1.0
	value = tmp.value

	t.rebalance(stack)

	return value, DeleteOk
}

func (t *WeightBalancedTree) Inspect() string {
	var out bytes.Buffer

	if t == nil {
		out.WriteString("nil")
	} else if t.isLeaf() {
		out.WriteString("(")
		out.WriteString("leaf/")
		out.WriteString(fmt.Sprintf("k:%v", t.key))
		out.WriteString("/")
		out.WriteString(fmt.Sprintf("w:%.2f", t.weight))
		out.WriteString("/")
		out.WriteString(fmt.Sprintf("v:%v", t.value))
		out.WriteString(")")
	} else {
		out.WriteString("[")
		out.WriteString(fmt.Sprintf("k:%v", t.key))
		out.WriteString("/")
		out.WriteString(fmt.Sprintf("w:%.2f", t.weight))
		out.WriteString("/")
		out.WriteString(fmt.Sprintf("l:%s", t.left.Inspect()))
		out.WriteString("/")
		out.WriteString(fmt.Sprintf("r:%s", t.right.Inspect()))
		out.WriteString("]")
	}

	return out.String()
}

func (*WeightBalancedTree) rebalance(stack []*WeightBalancedTree) {
	var tmp *WeightBalancedTree
	for len(stack) > 0 {
		stack, tmp = stack[:len(stack)-1], stack[len(stack)-1]
		if tmp.isLeaf() {
			continue
		}
		tmp.weight = tmp.left.weight + tmp.right.weight
		if tmp.right.weight < Alpha*tmp.weight {
			if tmp.left.left.weight > (Alpha+Epsilon)*tmp.weight {
				tmp.rotateRight()
				tmp.right.weight = tmp.right.left.weight + tmp.right.right.weight
			} else {
				tmp.left.rotateLeft()
				tmp.rotateRight()
				tmp.right.weight = tmp.right.left.weight + tmp.right.right.weight
				tmp.left.weight = tmp.left.left.weight + tmp.left.right.weight
			}
		} else if tmp.left.weight < Alpha*tmp.weight {
			if tmp.right.right.weight > (Alpha+Epsilon)*tmp.weight {
				tmp.rotateLeft()
				tmp.left.weight = tmp.left.left.weight + tmp.left.right.weight
			} else {
				tmp.right.rotateRight()
				tmp.rotateLeft()
				tmp.right.weight = tmp.right.right.weight + tmp.right.left.weight
				tmp.left.weight = tmp.left.left.weight + tmp.left.right.weight
			}
		}
	}
}

func (t *WeightBalancedTree) isLeaf() bool {
	return t.left == nil
}

func (t *WeightBalancedTree) isEmpty() bool {
	return t.isLeaf() && t.value == nil
}

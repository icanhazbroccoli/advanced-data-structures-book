package search_tree

import (
	"reflect"
	"testing"
)

func TestBinaryTreeRotateLeft(t *testing.T) {
	tree := &BinaryTree{
		key: IntKey(2),
		left: &BinaryTree{
			key:   IntKey(1),
			value: 1,
		},
		right: &BinaryTree{
			key: IntKey(3),
			left: &BinaryTree{
				key:   IntKey(2),
				value: 2,
			},
			right: &BinaryTree{
				key:   IntKey(3),
				value: 3,
			},
		},
	}

	expected := &BinaryTree{
		key: IntKey(3),
		left: &BinaryTree{
			key: IntKey(2),
			left: &BinaryTree{
				key:   IntKey(1),
				value: 1,
			},
			right: &BinaryTree{
				key:   IntKey(2),
				value: 2,
			},
		},
		right: &BinaryTree{
			key:   IntKey(3),
			value: 3,
		},
	}

	tree.RotateLeft()

	if !reflect.DeepEqual(tree, expected) {
		t.Fatalf("unexpected rotate left state: got=%s, want=%s",
			tree.Traverse(), expected.Traverse())
	}
}

func TestBinaryTreeRotateRight(t *testing.T) {
	tree := &BinaryTree{
		key: IntKey(3),
		left: &BinaryTree{
			key: IntKey(2),
			left: &BinaryTree{
				key:   IntKey(1),
				value: 1,
			},
			right: &BinaryTree{
				key:   IntKey(2),
				value: 2,
			},
		},
		right: &BinaryTree{
			key:   IntKey(3),
			value: 3,
		},
	}

	expected := &BinaryTree{
		key: IntKey(2),
		left: &BinaryTree{
			key:   IntKey(1),
			value: 1,
		},
		right: &BinaryTree{
			key: IntKey(3),
			left: &BinaryTree{
				key:   IntKey(2),
				value: 2,
			},
			right: &BinaryTree{
				key:   IntKey(3),
				value: 3,
			},
		},
	}

	tree.RotateRight()

	if !reflect.DeepEqual(tree, expected) {
		t.Fatalf("unexpected rotate right state: got=%s, want=%s",
			tree.Traverse(), expected.Traverse())
	}
}

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

func TestBinaryTreeInsert(t *testing.T) {
	tests := []struct {
		insertKey    int
		insertVal    int
		expectStatus bool
		expectTree   *BinaryTree
	}{
		{
			insertKey:    10,
			insertVal:    11,
			expectStatus: true,
			expectTree: &BinaryTree{
				key:   IntKey(10),
				value: 11,
			},
		},
		{
			insertKey:    20,
			insertVal:    22,
			expectStatus: true,
			expectTree: &BinaryTree{
				key: IntKey(20),
				left: &BinaryTree{
					key:   IntKey(10),
					value: 11,
				},
				right: &BinaryTree{
					key:   IntKey(20),
					value: 22,
				},
			},
		},
		{ // same as above: we expect no changes and InsertNone status
			insertKey:    20,
			insertVal:    23,
			expectStatus: false,
			expectTree: &BinaryTree{
				key: IntKey(20),
				left: &BinaryTree{
					key:   IntKey(10),
					value: 11,
				},
				right: &BinaryTree{
					key:   IntKey(20),
					value: 23,
				},
			},
		},
		{
			insertKey:    30,
			insertVal:    33,
			expectStatus: true,
			expectTree: &BinaryTree{
				key: IntKey(20),
				left: &BinaryTree{
					key:   IntKey(10),
					value: 11,
				},
				right: &BinaryTree{
					key: IntKey(30),
					left: &BinaryTree{
						key:   IntKey(20),
						value: 23,
					},
					right: &BinaryTree{
						key:   IntKey(30),
						value: 33,
					},
				},
			},
		},
		{
			insertKey:    5,
			insertVal:    55,
			expectStatus: true,
			expectTree: &BinaryTree{
				key: IntKey(20),
				left: &BinaryTree{
					key: IntKey(10),
					left: &BinaryTree{
						key:   IntKey(5),
						value: 55,
					},
					right: &BinaryTree{
						key:   IntKey(10),
						value: 11,
					},
				},
				right: &BinaryTree{
					key: IntKey(30),
					left: &BinaryTree{
						key:   IntKey(20),
						value: 23,
					},
					right: &BinaryTree{
						key:   IntKey(30),
						value: 33,
					},
				},
			},
		},
	}

	tree := NewBinaryTree()

	for _, tt := range tests {
		status := tree.Insert(IntKey(tt.insertKey), tt.insertVal)
		if status != InsertStatus(tt.expectStatus) {
			t.Errorf("unexpected insert status: got=%t, want=%t",
				status, tt.expectStatus)
		}

		if !reflect.DeepEqual(tree, tt.expectTree) {
			t.Errorf("unexpected tree state: got=%s, want: %s",
				tree.Traverse(), tt.expectTree.Traverse())
		}
	}
}

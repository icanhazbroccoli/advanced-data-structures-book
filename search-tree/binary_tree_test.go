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
			tree, expected)
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
			tree, expected)
	}
}

func TestBinaryTreeFind(t *testing.T) {
	tests := []struct {
		tree         *BinaryTree
		findKey      int
		expectVal    interface{}
		expectStatus bool
	}{
		{
			tree:         NewBinaryTree(),
			findKey:      42,
			expectStatus: false,
		},
		{
			tree: &BinaryTree{
				key:   IntKey(42),
				value: 43,
			},
			findKey:      42,
			expectStatus: true,
			expectVal:    43,
		},
		{
			tree: &BinaryTree{
				key:   IntKey(42),
				value: 43,
			},
			findKey:      122,
			expectStatus: false,
		},
		{
			tree: &BinaryTree{
				key: IntKey(2),
				left: &BinaryTree{
					key:   IntKey(1),
					value: 11,
				},
				right: &BinaryTree{
					key:   IntKey(2),
					value: 22,
				},
			},
			findKey:      2,
			expectStatus: true,
			expectVal:    22,
		},
		{
			tree: &BinaryTree{
				key: IntKey(2),
				left: &BinaryTree{
					key:   IntKey(1),
					value: 11,
				},
				right: &BinaryTree{
					key:   IntKey(2),
					value: 22,
				},
			},
			findKey:      1,
			expectStatus: true,
			expectVal:    11,
		},
		{
			tree: &BinaryTree{
				key: IntKey(2),
				left: &BinaryTree{
					key:   IntKey(1),
					value: 11,
				},
				right: &BinaryTree{
					key:   IntKey(2),
					value: 22,
				},
			},
			findKey:      3,
			expectStatus: false,
		},
		{
			tree: &BinaryTree{
				key: IntKey(3),
				left: &BinaryTree{
					key: IntKey(2),
					left: &BinaryTree{
						key:   IntKey(1),
						value: 11,
					},
					right: &BinaryTree{
						key:   IntKey(2),
						value: 22,
					},
				},
				right: &BinaryTree{
					key:   IntKey(3),
					value: 33,
				},
			},
			findKey:      1,
			expectStatus: true,
			expectVal:    11,
		},
		{
			tree: &BinaryTree{
				key: IntKey(3),
				left: &BinaryTree{
					key: IntKey(2),
					left: &BinaryTree{
						key:   IntKey(1),
						value: 11,
					},
					right: &BinaryTree{
						key:   IntKey(2),
						value: 22,
					},
				},
				right: &BinaryTree{
					key:   IntKey(3),
					value: 33,
				},
			},
			findKey:      2,
			expectStatus: true,
			expectVal:    22,
		},
		{
			tree: &BinaryTree{
				key: IntKey(3),
				left: &BinaryTree{
					key: IntKey(2),
					left: &BinaryTree{
						key:   IntKey(1),
						value: 11,
					},
					right: &BinaryTree{
						key:   IntKey(2),
						value: 22,
					},
				},
				right: &BinaryTree{
					key:   IntKey(3),
					value: 33,
				},
			},
			findKey:      2,
			expectStatus: true,
			expectVal:    22,
		},
		{
			tree: &BinaryTree{
				key: IntKey(3),
				left: &BinaryTree{
					key: IntKey(2),
					left: &BinaryTree{
						key:   IntKey(1),
						value: 11,
					},
					right: &BinaryTree{
						key:   IntKey(2),
						value: 22,
					},
				},
				right: &BinaryTree{
					key:   IntKey(3),
					value: 33,
				},
			},
			findKey:      3,
			expectStatus: true,
			expectVal:    33,
		},
		{
			tree: &BinaryTree{
				key: IntKey(3),
				left: &BinaryTree{
					key: IntKey(2),
					left: &BinaryTree{
						key:   IntKey(1),
						value: 11,
					},
					right: &BinaryTree{
						key:   IntKey(2),
						value: 22,
					},
				},
				right: &BinaryTree{
					key:   IntKey(3),
					value: 33,
				},
			},
			findKey:      4,
			expectStatus: false,
		},
	}

	for _, tt := range tests {
		value, ok := tt.tree.Find(IntKey(tt.findKey))
		if ok != FindStatus(tt.expectStatus) {
			t.Errorf("unexpected find status: got=%t, want=%t",
				ok, tt.expectStatus)
		}
		if ok == FindNone {
			continue
		}
		if !reflect.DeepEqual(value, tt.expectVal) {
			t.Errorf("unecpected find value: got=%v, want=%v",
				value, tt.expectVal)
		}
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
				tree, tt.expectTree)
		}
	}
}

func TestBinaryTreeDelete(t *testing.T) {
	tests := []struct {
		tree         *BinaryTree
		deleteKey    int
		expectStatus bool
		expectVal    interface{}
		expectTree   *BinaryTree
	}{
		{
			tree:         NewBinaryTree(),
			deleteKey:    42,
			expectStatus: false,
			expectTree:   NewBinaryTree(),
		},
		{
			tree: &BinaryTree{
				key:   IntKey(1),
				value: 11,
			},
			deleteKey:    1,
			expectStatus: true,
			expectTree:   NewBinaryTree(),
			expectVal:    11,
		},
		{
			tree: &BinaryTree{
				key:   IntKey(1),
				value: 11,
			},
			deleteKey:    2,
			expectStatus: false,
			expectTree: &BinaryTree{
				key:   IntKey(1),
				value: 11,
			},
		},
		{
			tree: &BinaryTree{
				key: IntKey(2),
				left: &BinaryTree{
					key:   IntKey(1),
					value: 11,
				},
				right: &BinaryTree{
					key:   IntKey(2),
					value: 22,
				},
			},
			deleteKey:    1,
			expectStatus: true,
			expectTree: &BinaryTree{
				key:   IntKey(2),
				value: 22,
			},
			expectVal: 11,
		},
		{
			tree: &BinaryTree{
				key: IntKey(2),
				left: &BinaryTree{
					key:   IntKey(1),
					value: 11,
				},
				right: &BinaryTree{
					key:   IntKey(2),
					value: 22,
				},
			},
			deleteKey:    2,
			expectStatus: true,
			expectTree: &BinaryTree{
				key:   IntKey(1),
				value: 11,
			},
			expectVal: 22,
		},
		{
			tree: &BinaryTree{
				key: IntKey(2),
				left: &BinaryTree{
					key:   IntKey(1),
					value: 11,
				},
				right: &BinaryTree{
					key: IntKey(3),
					left: &BinaryTree{
						key:   IntKey(2),
						value: 22,
					},
					right: &BinaryTree{
						key:   IntKey(3),
						value: 33,
					},
				},
			},
			deleteKey:    2,
			expectStatus: true,
			expectTree: &BinaryTree{
				key: IntKey(2),
				left: &BinaryTree{
					key:   IntKey(1),
					value: 11,
				},
				right: &BinaryTree{
					key:   IntKey(3),
					value: 33,
				},
			},
			expectVal: 22,
		},
		{
			tree: &BinaryTree{
				key: IntKey(2),
				left: &BinaryTree{
					key:   IntKey(1),
					value: 11,
				},
				right: &BinaryTree{
					key: IntKey(3),
					left: &BinaryTree{
						key:   IntKey(2),
						value: 22,
					},
					right: &BinaryTree{
						key:   IntKey(3),
						value: 33,
					},
				},
			},
			deleteKey:    3,
			expectStatus: true,
			expectTree: &BinaryTree{
				key: IntKey(2),
				left: &BinaryTree{
					key:   IntKey(1),
					value: 11,
				},
				right: &BinaryTree{
					key:   IntKey(2),
					value: 22,
				},
			},
			expectVal: 33,
		},
		{
			tree: &BinaryTree{
				key: IntKey(2),
				left: &BinaryTree{
					key:   IntKey(1),
					value: 11,
				},
				right: &BinaryTree{
					key: IntKey(3),
					left: &BinaryTree{
						key:   IntKey(2),
						value: 22,
					},
					right: &BinaryTree{
						key:   IntKey(3),
						value: 33,
					},
				},
			},
			deleteKey:    1,
			expectStatus: true,
			expectTree: &BinaryTree{
				key: IntKey(3),
				left: &BinaryTree{
					key:   IntKey(2),
					value: 22,
				},
				right: &BinaryTree{
					key:   IntKey(3),
					value: 33,
				},
			},
			expectVal: 11,
		},
	}

	for _, tt := range tests {
		value, ok := tt.tree.Delete(IntKey(tt.deleteKey))
		if ok != DeleteStatus(tt.expectStatus) {
			t.Errorf("unexpected delete status: got=%t, want=%t",
				ok, tt.expectStatus)
		}

		if !reflect.DeepEqual(tt.tree, tt.expectTree) {
			t.Errorf("unexpected tree state: got=%s, want=%s",
				tt.tree, tt.expectTree)
		}

		if ok != DeleteOk {
			continue
		}

		if !reflect.DeepEqual(value, tt.expectVal) {
			t.Errorf("unexpected delete value: got=%v, want=%v",
				value, tt.expectVal)
		}
	}
}

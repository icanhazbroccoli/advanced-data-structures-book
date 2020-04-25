package search_tree

import (
	"reflect"
	"testing"
)

func TestAVLTreeInsert(t *testing.T) {
	tests := []struct {
		insertKey    int
		insertVal    int
		expectStatus bool
		expectTree   *AVLTree
	}{
		{
			insertKey:    10,
			insertVal:    10,
			expectStatus: true,
			expectTree: &AVLTree{
				key:    IntKey(10),
				value:  10,
				height: 0,
			},
		},
		{
			insertKey:    20,
			insertVal:    20,
			expectStatus: true,
			expectTree: &AVLTree{
				key:    IntKey(20),
				height: 1,
				left: &AVLTree{
					key:    IntKey(10),
					value:  10,
					height: 0,
				},
				right: &AVLTree{
					key:    IntKey(20),
					value:  20,
					height: 0,
				},
			},
		},
		{
			insertKey:    15,
			insertVal:    15,
			expectStatus: true,
			expectTree: &AVLTree{
				key:    IntKey(20),
				height: 2,
				left: &AVLTree{
					key:    IntKey(15),
					height: 1,
					left: &AVLTree{
						key:    IntKey(10),
						value:  10,
						height: 0,
					},
					right: &AVLTree{
						key:    IntKey(15),
						value:  15,
						height: 0,
					},
				},
				right: &AVLTree{
					key:    IntKey(20),
					value:  20,
					height: 0,
				},
			},
		},
		{
			insertKey:    5,
			insertVal:    5,
			expectStatus: true,
			expectTree: &AVLTree{
				key:    IntKey(15),
				height: 2,
				left: &AVLTree{
					key:    IntKey(10),
					height: 1,
					left: &AVLTree{
						key:    IntKey(5),
						value:  5,
						height: 0,
					},
					right: &AVLTree{
						key:    IntKey(10),
						value:  10,
						height: 0,
					},
				},
				right: &AVLTree{
					key:    IntKey(20),
					height: 1,
					left: &AVLTree{
						key:    IntKey(15),
						value:  15,
						height: 0,
					},
					right: &AVLTree{
						key:    IntKey(20),
						value:  20,
						height: 0,
					},
				},
			},
		},
		{
			insertKey:    50,
			insertVal:    50,
			expectStatus: true,
			expectTree: &AVLTree{
				key:    IntKey(15),
				height: 3,
				left: &AVLTree{
					key:    IntKey(10),
					height: 1,
					left: &AVLTree{
						key:    IntKey(5),
						value:  5,
						height: 0,
					},
					right: &AVLTree{
						key:    IntKey(10),
						value:  10,
						height: 0,
					},
				},
				right: &AVLTree{
					key:    IntKey(20),
					height: 2,
					left: &AVLTree{
						key:    IntKey(15),
						value:  15,
						height: 0,
					},
					right: &AVLTree{
						key:    IntKey(50),
						height: 1,
						left: &AVLTree{
							key:    IntKey(20),
							value:  20,
							height: 0,
						},
						right: &AVLTree{
							key:    IntKey(50),
							value:  50,
							height: 0,
						},
					},
				},
			},
		},
	}

	tree := NewAVLTree()

	for _, tt := range tests {
		status := tree.Insert(IntKey(tt.insertKey), tt.insertVal)
		if status != InsertStatus(tt.expectStatus) {
			t.Errorf("unexpected insert status: got=%t, want=%t",
				status, tt.expectStatus)
		}

		if !reflect.DeepEqual(tree, tt.expectTree) {
			t.Errorf("unexpected tree state:\ngot=  %s\nwant= %s",
				tree, tt.expectTree)
		}
	}
}

func TestAVLTreeDelete(t *testing.T) {
	tests := []struct {
		tree         *AVLTree
		deleteKey    int
		expectStatus bool
		expectVal    interface{}
		expectTree   *AVLTree
	}{
		{
			tree: &AVLTree{
				key:    IntKey(15),
				height: 3,
				left: &AVLTree{
					key:    IntKey(10),
					height: 1,
					left: &AVLTree{
						key:    IntKey(5),
						value:  5,
						height: 0,
					},
					right: &AVLTree{
						key:    IntKey(10),
						value:  10,
						height: 0,
					},
				},
				right: &AVLTree{
					key:    IntKey(20),
					height: 2,
					left: &AVLTree{
						key:    IntKey(15),
						value:  15,
						height: 0,
					},
					right: &AVLTree{
						key:    IntKey(50),
						height: 1,
						left: &AVLTree{
							key:    IntKey(20),
							value:  20,
							height: 0,
						},
						right: &AVLTree{
							key:    IntKey(50),
							value:  50,
							height: 0,
						},
					},
				},
			},
			deleteKey:    50,
			expectStatus: true,
			expectVal:    50,
			expectTree: &AVLTree{
				key:    IntKey(15),
				height: 2,
				left: &AVLTree{
					key:    IntKey(10),
					height: 1,
					left: &AVLTree{
						key:    IntKey(5),
						value:  5,
						height: 0,
					},
					right: &AVLTree{
						key:    IntKey(10),
						value:  10,
						height: 0,
					},
				},
				right: &AVLTree{
					key:    IntKey(20),
					height: 1,
					left: &AVLTree{
						key:    IntKey(15),
						value:  15,
						height: 0,
					},
					right: &AVLTree{
						key:    IntKey(20),
						value:  20,
						height: 0,
					},
				},
			},
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

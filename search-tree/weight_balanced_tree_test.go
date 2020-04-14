package search_tree

import (
	"reflect"
	"testing"
)

func TestWeightBalancedTreeInsert(t *testing.T) {
	tests := []struct {
		insertKey    int
		insertVal    int
		expectStatus bool
		expectTree   *WeightBalancedTree
	}{
		{
			insertKey:    10,
			insertVal:    10,
			expectStatus: true,
			expectTree: &WeightBalancedTree{
				key:    IntKey(10),
				value:  10,
				weight: 1.0,
			},
		},
		{
			insertKey:    20,
			insertVal:    20,
			expectStatus: true,
			expectTree: &WeightBalancedTree{
				key:    IntKey(20),
				weight: 2.0,
				left: &WeightBalancedTree{
					key:    IntKey(10),
					value:  10,
					weight: 1.0,
				},
				right: &WeightBalancedTree{
					key:    IntKey(20),
					value:  20,
					weight: 1.0,
				},
			},
		},
		{
			insertKey:    15,
			insertVal:    15,
			expectStatus: true,
			expectTree: &WeightBalancedTree{
				key:    IntKey(20),
				weight: 3.0,
				left: &WeightBalancedTree{
					key:    IntKey(15),
					weight: 2.0,
					left: &WeightBalancedTree{
						key:    IntKey(10),
						value:  10,
						weight: 1.0,
					},
					right: &WeightBalancedTree{
						key:    IntKey(15),
						value:  15,
						weight: 1.0,
					},
				},
				right: &WeightBalancedTree{
					key:    IntKey(20),
					value:  20,
					weight: 1.0,
				},
			},
		},
		{
			insertKey:    5,
			insertVal:    5,
			expectStatus: true,
			expectTree: &WeightBalancedTree{
				key:    IntKey(15),
				weight: 4.0,
				left: &WeightBalancedTree{
					key:    IntKey(10),
					weight: 2.0,
					left: &WeightBalancedTree{
						key:    IntKey(5),
						value:  5,
						weight: 1.0,
					},
					right: &WeightBalancedTree{
						key:    IntKey(10),
						value:  10,
						weight: 1.0,
					},
				},
				right: &WeightBalancedTree{
					key:    IntKey(20),
					weight: 2.0,
					left: &WeightBalancedTree{
						key:    IntKey(15),
						value:  15,
						weight: 1.0,
					},
					right: &WeightBalancedTree{
						key:    IntKey(20),
						value:  20,
						weight: 1.0,
					},
				},
			},
		},
		{
			insertKey:    50,
			insertVal:    50,
			expectStatus: true,
			expectTree: &WeightBalancedTree{
				key:    IntKey(15),
				weight: 5.0,
				left: &WeightBalancedTree{
					key:    IntKey(10),
					weight: 2.0,
					left: &WeightBalancedTree{
						key:    IntKey(5),
						value:  5,
						weight: 1.0,
					},
					right: &WeightBalancedTree{
						key:    IntKey(10),
						value:  10,
						weight: 1.0,
					},
				},
				right: &WeightBalancedTree{
					key:    IntKey(20),
					weight: 3.0,
					left: &WeightBalancedTree{
						key:    IntKey(15),
						value:  15,
						weight: 1.0,
					},
					right: &WeightBalancedTree{
						key:    IntKey(50),
						weight: 2.0,
						left: &WeightBalancedTree{
							key:    IntKey(20),
							value:  20,
							weight: 1.0,
						},
						right: &WeightBalancedTree{
							key:    IntKey(50),
							value:  50,
							weight: 1.0,
						},
					},
				},
			},
		},
	}

	tree := NewWeightBalancedTree()

	for _, tt := range tests {
		status := tree.Insert(IntKey(tt.insertKey), tt.insertVal)
		if status != InsertStatus(tt.expectStatus) {
			t.Errorf("unexpected insert status: got=%t, want=%t",
				status, tt.expectStatus)
		}

		if !reflect.DeepEqual(tree, tt.expectTree) {
			t.Errorf("unexpected tree state:\ngot=  %s\nwant= %s",
				tree.Traverse(), tt.expectTree.Traverse())
		}
	}
}

func TestWeightBalancedTreeDelete(t *testing.T) {
	tests := []struct {
		tree         *WeightBalancedTree
		deleteKey    int
		expectStatus bool
		expectVal    interface{}
		expectTree   *WeightBalancedTree
	}{
		{
			tree: &WeightBalancedTree{
				key:    IntKey(15),
				weight: 5.0,
				left: &WeightBalancedTree{
					key:    IntKey(10),
					weight: 2.0,
					left: &WeightBalancedTree{
						key:    IntKey(5),
						value:  5,
						weight: 1.0,
					},
					right: &WeightBalancedTree{
						key:    IntKey(10),
						value:  10,
						weight: 1.0,
					},
				},
				right: &WeightBalancedTree{
					key:    IntKey(20),
					weight: 3.0,
					left: &WeightBalancedTree{
						key:    IntKey(15),
						value:  15,
						weight: 1.0,
					},
					right: &WeightBalancedTree{
						key:    IntKey(50),
						weight: 2.0,
						left: &WeightBalancedTree{
							key:    IntKey(20),
							value:  20,
							weight: 1.0,
						},
						right: &WeightBalancedTree{
							key:    IntKey(50),
							value:  50,
							weight: 1.0,
						},
					},
				},
			},
			deleteKey:    50,
			expectStatus: true,
			expectVal:    50,
			expectTree: &WeightBalancedTree{
				key:    IntKey(15),
				weight: 4.0,
				left: &WeightBalancedTree{
					key:    IntKey(10),
					weight: 2.0,
					left: &WeightBalancedTree{
						key:    IntKey(5),
						value:  5,
						weight: 1.0,
					},
					right: &WeightBalancedTree{
						key:    IntKey(10),
						value:  10,
						weight: 1.0,
					},
				},
				right: &WeightBalancedTree{
					key:    IntKey(20),
					weight: 2.0,
					left: &WeightBalancedTree{
						key:    IntKey(15),
						value:  15,
						weight: 1.0,
					},
					right: &WeightBalancedTree{
						key:    IntKey(20),
						value:  20,
						weight: 1.0,
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
				tt.tree.Traverse(), tt.expectTree.Traverse())
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

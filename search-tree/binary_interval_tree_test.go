package search_tree

import (
	"reflect"
	"testing"
)

func TestBinaryTreeFindInterval(t *testing.T) {
	tests := []struct {
		tree *BinaryTree
		a, b int
		exp  []int
	}{
		{
			tree: &BinaryTree{
				key: IntKey(20),
				left: &BinaryTree{
					key:   IntKey(10),
					value: 11,
				},
				right: &BinaryTree{
					key: IntKey(30),
					left: &BinaryTree{
						key:   IntKey(20),
						value: 22,
					},
					right: &BinaryTree{
						key:   IntKey(30),
						value: 33,
					},
				},
			},
			a:   10,
			b:   20,
			exp: []int{11},
		},
		{
			tree: &BinaryTree{
				key: IntKey(20),
				left: &BinaryTree{
					key:   IntKey(10),
					value: 11,
				},
				right: &BinaryTree{
					key: IntKey(30),
					left: &BinaryTree{
						key:   IntKey(20),
						value: 22,
					},
					right: &BinaryTree{
						key:   IntKey(30),
						value: 33,
					},
				},
			},
			a:   10,
			b:   25,
			exp: []int{11, 22},
		},
		{
			tree: &BinaryTree{
				key: IntKey(20),
				left: &BinaryTree{
					key:   IntKey(10),
					value: 11,
				},
				right: &BinaryTree{
					key: IntKey(30),
					left: &BinaryTree{
						key:   IntKey(20),
						value: 22,
					},
					right: &BinaryTree{
						key:   IntKey(30),
						value: 33,
					},
				},
			},
			a:   15,
			b:   20,
			exp: []int{},
		},
		{
			tree: &BinaryTree{
				key: IntKey(20),
				left: &BinaryTree{
					key:   IntKey(10),
					value: 11,
				},
				right: &BinaryTree{
					key: IntKey(30),
					left: &BinaryTree{
						key:   IntKey(20),
						value: 22,
					},
					right: &BinaryTree{
						key:   IntKey(30),
						value: 33,
					},
				},
			},
			a:   20,
			b:   30,
			exp: []int{22},
		},
		{
			tree: &BinaryTree{
				key: IntKey(20),
				left: &BinaryTree{
					key:   IntKey(10),
					value: 11,
				},
				right: &BinaryTree{
					key: IntKey(30),
					left: &BinaryTree{
						key:   IntKey(20),
						value: 22,
					},
					right: &BinaryTree{
						key:   IntKey(30),
						value: 33,
					},
				},
			},
			a:   20,
			b:   35,
			exp: []int{22, 33},
		},
		{
			tree: &BinaryTree{
				key: IntKey(20),
				left: &BinaryTree{
					key:   IntKey(10),
					value: 11,
				},
				right: &BinaryTree{
					key: IntKey(30),
					left: &BinaryTree{
						key:   IntKey(20),
						value: 22,
					},
					right: &BinaryTree{
						key:   IntKey(30),
						value: 33,
					},
				},
			},
			a:   40,
			b:   50,
			exp: []int{},
		},
		{
			tree: &BinaryTree{
				key: IntKey(20),
				left: &BinaryTree{
					key:   IntKey(10),
					value: 11,
				},
				right: &BinaryTree{
					key: IntKey(30),
					left: &BinaryTree{
						key:   IntKey(20),
						value: 22,
					},
					right: &BinaryTree{
						key:   IntKey(30),
						value: 33,
					},
				},
			},
			a:   20,
			b:   10,
			exp: []int{},
		},
	}

	for _, tt := range tests {
		tree := &BinaryIntervalTree{tt.tree}
		exp := make([]StoredObject, 0, len(tt.exp))
		for _, e := range tt.exp {
			exp = append(exp, e)
		}
		res := tree.FindInterval(IntKey(tt.a), IntKey(tt.b))
		if !reflect.DeepEqual(res, exp) {
			t.Errorf("unexpected interval find result on tree: %s: got=%+v, want=%+v",
				tree.Traverse(), res, exp)
		}
	}
}

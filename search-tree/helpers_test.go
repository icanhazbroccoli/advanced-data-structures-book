package search_tree

import (
	"reflect"
	"testing"
)

func TestMakeTreeBottomUp(t *testing.T) {
	tests := []struct {
		input    *BinaryTree
		expected *BinaryTree
	}{
		{
			input:    nil,
			expected: NewBinaryTree(),
		},
		{
			input: &BinaryTree{
				key:   IntKey(10),
				value: 11,
			},
			expected: &BinaryTree{
				key:   IntKey(10),
				value: 11,
			},
		},
		{
			input: &BinaryTree{
				key:   IntKey(1),
				value: 11,
				right: &BinaryTree{
					key:   IntKey(2),
					value: 22,
					right: &BinaryTree{
						key:   IntKey(3),
						value: 33,
						right: &BinaryTree{
							key:   IntKey(4),
							value: 44,
							right: &BinaryTree{
								key:   IntKey(5),
								value: 55,
								right: &BinaryTree{
									key:   IntKey(6),
									value: 66,
									right: &BinaryTree{
										key:   IntKey(7),
										value: 77,
									},
								},
							},
						},
					},
				},
			},
			expected: &BinaryTree{
				key: IntKey(5),
				left: &BinaryTree{
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
						key: IntKey(4),
						left: &BinaryTree{
							key:   IntKey(3),
							value: 33,
						},
						right: &BinaryTree{
							key:   IntKey(4),
							value: 44,
						},
					},
				},
				right: &BinaryTree{
					key: IntKey(7),
					left: &BinaryTree{
						key: IntKey(6),
						left: &BinaryTree{
							key:   IntKey(5),
							value: 55,
						},
						right: &BinaryTree{
							key:   IntKey(6),
							value: 66,
						},
					},
					right: &BinaryTree{
						key:   IntKey(7),
						value: 77,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		res := MakeTreeBottomUp(tt.input)
		if !reflect.DeepEqual(res, tt.expected) {
			t.Errorf("unexpected tree structure:\ngot= %s\nwant=%s",
				res.Traverse(), tt.expected.Traverse())
		}
	}
}

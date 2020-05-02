package search_tree

import (
	"reflect"
	"testing"
)

func TestABTreeFind(t *testing.T) {
	tree := &ABTree{
		A:      4,
		B:      8,
		degree: 4,
		height: 1,
		keys:   []SearchKey{nil, IntKey(10), IntKey(20), IntKey(50)},
		next: []*ABTree{
			&ABTree{
				A:      4,
				B:      8,
				degree: 5,
				height: 0,
				keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7),
					IntKey(9)},
				values: []StoredObject{1, 4, 6, 7, 9},
			},
			&ABTree{
				A:      4,
				B:      8,
				degree: 7,
				height: 0,
				keys: []SearchKey{IntKey(10), IntKey(12), IntKey(13), IntKey(14),
					IntKey(15), IntKey(17), IntKey(19)},
				values: []StoredObject{10, 12, 13, 14, 15, 17, 19},
			},
			&ABTree{
				A:      4,
				B:      8,
				degree: 4,
				height: 0,
				keys:   []SearchKey{IntKey(21), IntKey(23), IntKey(24), IntKey(42)},
				values: []StoredObject{21, 23, 24, 42},
			},
			&ABTree{
				A:      4,
				B:      8,
				degree: 4,
				height: 0,
				keys:   []SearchKey{IntKey(55), IntKey(62), IntKey(70), IntKey(88)},
				values: []StoredObject{55, 62, 70, 88},
			},
		},
	}

	tests := []struct {
		key       SearchKey
		expStatus FindStatus
		expVal    StoredObject
	}{
		{key: IntKey(0), expStatus: FindNone},
		{key: IntKey(1), expStatus: FindOk, expVal: 1},
		{key: IntKey(3), expStatus: FindNone},
		{key: IntKey(4), expStatus: FindOk, expVal: 4},
		{key: IntKey(5), expStatus: FindNone},
		{key: IntKey(6), expStatus: FindOk, expVal: 6},
		{key: IntKey(7), expStatus: FindOk, expVal: 7},
		{key: IntKey(8), expStatus: FindNone},
		{key: IntKey(9), expStatus: FindOk, expVal: 9},
		{key: IntKey(10), expStatus: FindOk, expVal: 10},
		{key: IntKey(11), expStatus: FindNone},
		{key: IntKey(12), expStatus: FindOk, expVal: 12},
		{key: IntKey(13), expStatus: FindOk, expVal: 13},
		{key: IntKey(14), expStatus: FindOk, expVal: 14},
		{key: IntKey(15), expStatus: FindOk, expVal: 15},
		{key: IntKey(16), expStatus: FindNone},
		{key: IntKey(17), expStatus: FindOk, expVal: 17},
		{key: IntKey(18), expStatus: FindNone},
		{key: IntKey(19), expStatus: FindOk, expVal: 19},
		{key: IntKey(20), expStatus: FindNone},
		{key: IntKey(21), expStatus: FindOk, expVal: 21},
		{key: IntKey(22), expStatus: FindNone},
		{key: IntKey(23), expStatus: FindOk, expVal: 23},
		{key: IntKey(24), expStatus: FindOk, expVal: 24},
		{key: IntKey(25), expStatus: FindNone},
		{key: IntKey(42), expStatus: FindOk, expVal: 42},
		{key: IntKey(50), expStatus: FindNone},
		{key: IntKey(55), expStatus: FindOk, expVal: 55},
		{key: IntKey(60), expStatus: FindNone},
		{key: IntKey(62), expStatus: FindOk, expVal: 62},
		{key: IntKey(65), expStatus: FindNone},
		{key: IntKey(70), expStatus: FindOk, expVal: 70},
		{key: IntKey(71), expStatus: FindNone},
		{key: IntKey(88), expStatus: FindOk, expVal: 88},
		{key: IntKey(90), expStatus: FindNone},
	}

	for _, tt := range tests {
		val, ok := tree.Find(tt.key)
		if ok != tt.expStatus {
			t.Errorf("Unexpected FindStatus for key %q: got=%t want=%t",
				tt.key, ok, tt.expStatus)
		}

		if ok == FindOk {
			if !reflect.DeepEqual(val, tt.expVal) {
				t.Errorf("Unexpected StoredObject for key %q: got=%+v want=%+v",
					tt.key, val, tt.expVal)
			}
		}
	}
}

func TestABTreeInsert(t *testing.T) {
	A, B := 4, 8
	tests := []struct {
		key      SearchKey
		value    StoredObject
		wantTree *ABTree
	}{
		{
			key:   IntKey(1),
			value: 11,
			wantTree: &ABTree{
				A:      4,
				B:      8,
				degree: 1,
				height: 0,
				keys:   []SearchKey{IntKey(1)},
				next:   []*ABTree{nil},
				values: []StoredObject{11},
			},
		},
		{ // same key as above, checking duplicate keys
			key:   IntKey(1),
			value: 1, // the value has been changed to ensure the update
			wantTree: &ABTree{
				A:      4,
				B:      8,
				degree: 1,
				height: 0,
				keys:   []SearchKey{IntKey(1)},
				next:   []*ABTree{nil},
				values: []StoredObject{1},
			},
		},
		{
			key:   IntKey(4),
			value: 4,
			wantTree: &ABTree{
				A:      4,
				B:      8,
				degree: 2,
				height: 0,
				keys:   []SearchKey{IntKey(1), IntKey(4)},
				next:   []*ABTree{nil, nil},
				values: []StoredObject{1, 4},
			},
		},
		{
			key:   IntKey(6),
			value: 6,
			wantTree: &ABTree{
				A:      4,
				B:      8,
				degree: 3,
				height: 0,
				keys:   []SearchKey{IntKey(1), IntKey(4), IntKey(6)},
				next:   []*ABTree{nil, nil, nil},
				values: []StoredObject{1, 4, 6},
			},
		},
		{
			key:   IntKey(7),
			value: 7,
			wantTree: &ABTree{
				A:      4,
				B:      8,
				degree: 4,
				height: 0,
				keys:   []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7)},
				next:   []*ABTree{nil, nil, nil, nil},
				values: []StoredObject{1, 4, 6, 7},
			},
		},
		{
			key:   IntKey(9),
			value: 9,
			wantTree: &ABTree{
				A:      4,
				B:      8,
				degree: 5,
				height: 0,
				keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7),
					IntKey(9)},
				next:   []*ABTree{nil, nil, nil, nil, nil},
				values: []StoredObject{1, 4, 6, 7, 9},
			},
		},
		{
			key:   IntKey(10),
			value: 10,
			wantTree: &ABTree{
				A:      4,
				B:      8,
				degree: 6,
				height: 0,
				keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7),
					IntKey(9), IntKey(10)},
				next:   []*ABTree{nil, nil, nil, nil, nil, nil},
				values: []StoredObject{1, 4, 6, 7, 9, 10},
			},
		},
		{
			key:   IntKey(12),
			value: 12,
			wantTree: &ABTree{
				A:      4,
				B:      8,
				degree: 7,
				height: 0,
				keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7),
					IntKey(9), IntKey(10), IntKey(12)},
				next:   []*ABTree{nil, nil, nil, nil, nil, nil, nil},
				values: []StoredObject{1, 4, 6, 7, 9, 10, 12},
			},
		},
		{
			key:   IntKey(13),
			value: 13,
			wantTree: &ABTree{
				A:      4,
				B:      8,
				degree: 8,
				height: 0,
				keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7),
					IntKey(9), IntKey(10), IntKey(12), IntKey(13)},
				next:   []*ABTree{nil, nil, nil, nil, nil, nil, nil, nil},
				values: []StoredObject{1, 4, 6, 7, 9, 10, 12, 13},
			},
		},
		{
			key:   IntKey(14),
			value: 14,
			wantTree: &ABTree{
				A:      4,
				B:      8,
				degree: 2,
				height: 1,
				keys:   []SearchKey{nil, IntKey(10)},
				next: []*ABTree{
					&ABTree{
						A:      4,
						B:      8,
						degree: 5,
						height: 0,
						keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7),
							IntKey(9)},
						values: []StoredObject{1, 4, 6, 7, 9},
						next:   []*ABTree{nil, nil, nil, nil, nil},
					},
					&ABTree{
						A:      4,
						B:      8,
						degree: 4,
						height: 0,
						keys:   []SearchKey{IntKey(10), IntKey(12), IntKey(13), IntKey(14)},
						values: []StoredObject{10, 12, 13, 14},
						next:   []*ABTree{nil, nil, nil, nil},
					}},
				values: []StoredObject{nil, nil},
			},
		},
		{
			key:   IntKey(15),
			value: 15,
			wantTree: &ABTree{
				A:      4,
				B:      8,
				degree: 2,
				height: 1,
				keys:   []SearchKey{nil, IntKey(10)},
				next: []*ABTree{
					&ABTree{
						A:      4,
						B:      8,
						degree: 5,
						height: 0,
						keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7),
							IntKey(9)},
						values: []StoredObject{1, 4, 6, 7, 9},
						next:   []*ABTree{nil, nil, nil, nil, nil},
					},
					&ABTree{
						A:      4,
						B:      8,
						degree: 5,
						height: 0,
						keys: []SearchKey{IntKey(10), IntKey(12), IntKey(13), IntKey(14),
							IntKey(15)},
						values: []StoredObject{10, 12, 13, 14, 15},
						next:   []*ABTree{nil, nil, nil, nil, nil},
					}},
				values: []StoredObject{nil, nil},
			},
		},
		{
			key:   IntKey(17),
			value: 17,
			wantTree: &ABTree{
				A:      4,
				B:      8,
				degree: 2,
				height: 1,
				keys:   []SearchKey{nil, IntKey(10)},
				next: []*ABTree{
					&ABTree{
						A:      4,
						B:      8,
						degree: 5,
						height: 0,
						keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7),
							IntKey(9)},
						values: []StoredObject{1, 4, 6, 7, 9},
						next:   []*ABTree{nil, nil, nil, nil, nil},
					},
					&ABTree{
						A:      4,
						B:      8,
						degree: 6,
						height: 0,
						keys: []SearchKey{IntKey(10), IntKey(12), IntKey(13), IntKey(14),
							IntKey(15), IntKey(17)},
						values: []StoredObject{10, 12, 13, 14, 15, 17},
						next:   []*ABTree{nil, nil, nil, nil, nil, nil},
					}},
				values: []StoredObject{nil, nil},
			},
		},
		{
			key:   IntKey(19),
			value: 19,
			wantTree: &ABTree{
				A:      4,
				B:      8,
				degree: 2,
				height: 1,
				keys:   []SearchKey{nil, IntKey(10)},
				next: []*ABTree{
					&ABTree{
						A:      4,
						B:      8,
						degree: 5,
						height: 0,
						keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7),
							IntKey(9)},
						values: []StoredObject{1, 4, 6, 7, 9},
						next:   []*ABTree{nil, nil, nil, nil, nil},
					},
					&ABTree{
						A:      4,
						B:      8,
						degree: 7,
						height: 0,
						keys: []SearchKey{IntKey(10), IntKey(12), IntKey(13), IntKey(14),
							IntKey(15), IntKey(17), IntKey(19)},
						values: []StoredObject{10, 12, 13, 14, 15, 17, 19},
						next:   []*ABTree{nil, nil, nil, nil, nil, nil, nil},
					}},
				values: []StoredObject{nil, nil},
			},
		},
		{
			key:   IntKey(21),
			value: 21,
			wantTree: &ABTree{
				A:      4,
				B:      8,
				degree: 2,
				height: 1,
				keys:   []SearchKey{nil, IntKey(10)},
				next: []*ABTree{
					&ABTree{
						A:      4,
						B:      8,
						degree: 5,
						height: 0,
						keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7),
							IntKey(9)},
						values: []StoredObject{1, 4, 6, 7, 9},
						next:   []*ABTree{nil, nil, nil, nil, nil},
					},
					&ABTree{
						A:      4,
						B:      8,
						degree: 8,
						height: 0,
						keys: []SearchKey{IntKey(10), IntKey(12), IntKey(13), IntKey(14),
							IntKey(15), IntKey(17), IntKey(19), IntKey(21)},
						values: []StoredObject{10, 12, 13, 14, 15, 17, 19, 21},
						next:   []*ABTree{nil, nil, nil, nil, nil, nil, nil, nil},
					}},
				values: []StoredObject{nil, nil},
			},
		},
		{
			key:   IntKey(23),
			value: 23,
			wantTree: &ABTree{
				A:      4,
				B:      8,
				degree: 3,
				height: 1,
				keys:   []SearchKey{nil, IntKey(10), IntKey(17)},
				next: []*ABTree{
					&ABTree{
						A:      4,
						B:      8,
						degree: 5,
						height: 0,
						keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7),
							IntKey(9)},
						values: []StoredObject{1, 4, 6, 7, 9},
						next:   []*ABTree{nil, nil, nil, nil, nil},
					},
					&ABTree{
						A:      4,
						B:      8,
						degree: 5,
						height: 0,
						keys: []SearchKey{IntKey(10), IntKey(12), IntKey(13), IntKey(14),
							IntKey(15)},
						values: []StoredObject{10, 12, 13, 14, 15},
						next:   []*ABTree{nil, nil, nil, nil, nil},
					},
					&ABTree{
						A:      4,
						B:      8,
						degree: 4,
						height: 0,
						keys:   []SearchKey{IntKey(17), IntKey(19), IntKey(21), IntKey(23)},
						values: []StoredObject{17, 19, 21, 23},
						next:   []*ABTree{nil, nil, nil, nil},
					}},
				values: []StoredObject{nil, nil, nil},
			},
		},
	}

	tree := NewABTree(A, B)

	for _, tt := range tests {
		tree.Insert(tt.key, tt.value)
		_ = tt.wantTree.String()
		if !reflect.DeepEqual(tree, tt.wantTree) {
			t.Fatalf("Unexpected state of the tree:\ngot: %s\nwant: %s",
				tree, tt.wantTree)
		}
	}
}

func TestABTreeDelete(t *testing.T) {
	tests := []struct {
		inputTree, expectTree *ABTree
		deleteKey             SearchKey
		expectObj             StoredObject
		expectStatus          DeleteStatus
	}{
		{
			inputTree:    &ABTree{A: 4, B: 8, degree: 0},
			deleteKey:    IntKey(42),
			expectObj:    nil,
			expectStatus: DeleteNone,
			expectTree:   &ABTree{A: 4, B: 8, degree: 0},
		},
		{
			inputTree: &ABTree{
				A:      4,
				B:      8,
				keys:   []SearchKey{IntKey(1)},
				next:   []*ABTree{nil},
				values: []StoredObject{1},
				degree: 1,
			},
			deleteKey:    IntKey(1),
			expectObj:    1,
			expectStatus: DeleteOk,
			expectTree: &ABTree{
				A:      4,
				B:      8,
				keys:   []SearchKey{},
				next:   []*ABTree{},
				values: []StoredObject{},
				degree: 0,
			},
		},
		{
			inputTree: &ABTree{
				A:      4,
				B:      8,
				keys:   []SearchKey{IntKey(1), IntKey(2)},
				next:   []*ABTree{nil, nil},
				values: []StoredObject{1, 2},
				degree: 2,
			},
			deleteKey:    IntKey(1),
			expectObj:    1,
			expectStatus: DeleteOk,
			expectTree: &ABTree{
				A:      4,
				B:      8,
				keys:   []SearchKey{IntKey(2)},
				next:   []*ABTree{nil},
				values: []StoredObject{2},
				degree: 1,
			},
		},
		{
			inputTree: &ABTree{
				A:      4,
				B:      8,
				keys:   []SearchKey{IntKey(1), IntKey(2)},
				next:   []*ABTree{nil, nil},
				values: []StoredObject{1, 2},
				degree: 2,
			},
			deleteKey:    IntKey(2),
			expectObj:    2,
			expectStatus: DeleteOk,
			expectTree: &ABTree{
				A:      4,
				B:      8,
				keys:   []SearchKey{IntKey(1)},
				next:   []*ABTree{nil},
				values: []StoredObject{1},
				degree: 1,
			},
		},
		{
			inputTree: &ABTree{
				A: 4,
				B: 8,
				keys: []SearchKey{IntKey(1), IntKey(2), IntKey(3), IntKey(4),
					IntKey(5), IntKey(6), IntKey(7), IntKey(8)},
				next:   []*ABTree{nil, nil, nil, nil, nil, nil, nil, nil},
				values: []StoredObject{1, 2, 3, 4, 5, 6, 7, 8},
				degree: 8,
			},
			deleteKey:    IntKey(1),
			expectObj:    1,
			expectStatus: DeleteOk,
			expectTree: &ABTree{
				A: 4,
				B: 8,
				keys: []SearchKey{IntKey(2), IntKey(3), IntKey(4), IntKey(5),
					IntKey(6), IntKey(7), IntKey(8)},
				next:   []*ABTree{nil, nil, nil, nil, nil, nil, nil},
				values: []StoredObject{2, 3, 4, 5, 6, 7, 8},
				degree: 7,
			},
		},
		{
			inputTree: &ABTree{
				A: 4,
				B: 8,
				keys: []SearchKey{IntKey(1), IntKey(2), IntKey(3), IntKey(4),
					IntKey(5), IntKey(6), IntKey(7), IntKey(8)},
				next:   []*ABTree{nil, nil, nil, nil, nil, nil, nil, nil},
				values: []StoredObject{1, 2, 3, 4, 5, 6, 7, 8},
				degree: 8,
			},
			deleteKey:    IntKey(5),
			expectObj:    5,
			expectStatus: DeleteOk,
			expectTree: &ABTree{
				A: 4,
				B: 8,
				keys: []SearchKey{IntKey(1), IntKey(2), IntKey(3), IntKey(4),
					IntKey(6), IntKey(7), IntKey(8)},
				next:   []*ABTree{nil, nil, nil, nil, nil, nil, nil},
				values: []StoredObject{1, 2, 3, 4, 6, 7, 8},
				degree: 7,
			},
		},
		{
			inputTree: &ABTree{
				A: 4,
				B: 8,
				keys: []SearchKey{IntKey(1), IntKey(2), IntKey(3), IntKey(4),
					IntKey(5), IntKey(6), IntKey(7), IntKey(8)},
				next:   []*ABTree{nil, nil, nil, nil, nil, nil, nil, nil},
				values: []StoredObject{1, 2, 3, 4, 5, 6, 7, 8},
				degree: 8,
			},
			deleteKey:    IntKey(8),
			expectObj:    8,
			expectStatus: DeleteOk,
			expectTree: &ABTree{
				A: 4,
				B: 8,
				keys: []SearchKey{IntKey(1), IntKey(2), IntKey(3), IntKey(4),
					IntKey(5), IntKey(6), IntKey(7)},
				next:   []*ABTree{nil, nil, nil, nil, nil, nil, nil},
				values: []StoredObject{1, 2, 3, 4, 5, 6, 7},
				degree: 7,
			},
		},
		{
			inputTree: &ABTree{
				A:      4,
				B:      8,
				degree: 2,
				height: 1,
				keys:   []SearchKey{nil, IntKey(10)},
				next: []*ABTree{
					&ABTree{
						A:      4,
						B:      8,
						degree: 5,
						height: 0,
						keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7),
							IntKey(9)},
						values: []StoredObject{1, 4, 6, 7, 9},
						next:   []*ABTree{nil, nil, nil, nil, nil},
					},
					&ABTree{
						A:      4,
						B:      8,
						degree: 4,
						height: 0,
						keys:   []SearchKey{IntKey(10), IntKey(12), IntKey(13), IntKey(14)},
						values: []StoredObject{10, 12, 13, 14},
						next:   []*ABTree{nil, nil, nil, nil},
					}},
				values: []StoredObject{nil, nil},
			},
			deleteKey:    IntKey(1),
			expectObj:    1,
			expectStatus: DeleteOk,
			expectTree: &ABTree{
				A:      4,
				B:      8,
				degree: 2,
				height: 1,
				keys:   []SearchKey{nil, IntKey(10)},
				next: []*ABTree{
					&ABTree{
						A:      4,
						B:      8,
						degree: 4,
						height: 0,
						keys: []SearchKey{IntKey(4), IntKey(6), IntKey(7),
							IntKey(9)},
						values: []StoredObject{4, 6, 7, 9},
						next:   []*ABTree{nil, nil, nil, nil},
					},
					&ABTree{
						A:      4,
						B:      8,
						degree: 4,
						height: 0,
						keys:   []SearchKey{IntKey(10), IntKey(12), IntKey(13), IntKey(14)},
						values: []StoredObject{10, 12, 13, 14},
						next:   []*ABTree{nil, nil, nil, nil},
					}},
				values: []StoredObject{nil, nil},
			},
		},
		{
			inputTree: &ABTree{
				A:      4,
				B:      8,
				degree: 2,
				height: 1,
				keys:   []SearchKey{nil, IntKey(10)},
				next: []*ABTree{
					&ABTree{
						A:      4,
						B:      8,
						degree: 5,
						height: 0,
						keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7),
							IntKey(9)},
						values: []StoredObject{1, 4, 6, 7, 9},
						next:   []*ABTree{nil, nil, nil, nil, nil},
					},
					&ABTree{
						A:      4,
						B:      8,
						degree: 4,
						height: 0,
						keys:   []SearchKey{IntKey(10), IntKey(12), IntKey(13), IntKey(14)},
						values: []StoredObject{10, 12, 13, 14},
						next:   []*ABTree{nil, nil, nil, nil},
					}},
				values: []StoredObject{nil, nil},
			},
			deleteKey:    IntKey(6),
			expectObj:    6,
			expectStatus: DeleteOk,
			expectTree: &ABTree{
				A:      4,
				B:      8,
				degree: 2,
				height: 1,
				keys:   []SearchKey{nil, IntKey(10)},
				next: []*ABTree{
					&ABTree{
						A:      4,
						B:      8,
						degree: 4,
						height: 0,
						keys:   []SearchKey{IntKey(1), IntKey(4), IntKey(7), IntKey(9)},
						values: []StoredObject{1, 4, 7, 9},
						next:   []*ABTree{nil, nil, nil, nil},
					},
					&ABTree{
						A:      4,
						B:      8,
						degree: 4,
						height: 0,
						keys:   []SearchKey{IntKey(10), IntKey(12), IntKey(13), IntKey(14)},
						values: []StoredObject{10, 12, 13, 14},
						next:   []*ABTree{nil, nil, nil, nil},
					}},
				values: []StoredObject{nil, nil},
			},
		},
		{
			inputTree: &ABTree{
				A:      4,
				B:      8,
				degree: 2,
				height: 1,
				keys:   []SearchKey{nil, IntKey(10)},
				next: []*ABTree{
					&ABTree{
						A:      4,
						B:      8,
						degree: 5,
						height: 0,
						keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7),
							IntKey(9)},
						values: []StoredObject{1, 4, 6, 7, 9},
						next:   []*ABTree{nil, nil, nil, nil, nil},
					},
					&ABTree{
						A:      4,
						B:      8,
						degree: 4,
						height: 0,
						keys:   []SearchKey{IntKey(10), IntKey(12), IntKey(13), IntKey(14)},
						values: []StoredObject{10, 12, 13, 14},
						next:   []*ABTree{nil, nil, nil, nil},
					}},
				values: []StoredObject{nil, nil},
			},
			deleteKey:    IntKey(9),
			expectObj:    9,
			expectStatus: DeleteOk,
			expectTree: &ABTree{
				A:      4,
				B:      8,
				degree: 2,
				height: 1,
				keys:   []SearchKey{nil, IntKey(10)},
				next: []*ABTree{
					&ABTree{
						A:      4,
						B:      8,
						degree: 4,
						height: 0,
						keys:   []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7)},
						values: []StoredObject{1, 4, 6, 7},
						next:   []*ABTree{nil, nil, nil, nil},
					},
					&ABTree{
						A:      4,
						B:      8,
						degree: 4,
						height: 0,
						keys:   []SearchKey{IntKey(10), IntKey(12), IntKey(13), IntKey(14)},
						values: []StoredObject{10, 12, 13, 14},
						next:   []*ABTree{nil, nil, nil, nil},
					}},
				values: []StoredObject{nil, nil},
			},
		},
		{
			inputTree: &ABTree{
				A:      4,
				B:      8,
				degree: 2,
				height: 1,
				keys:   []SearchKey{nil, IntKey(10)},
				next: []*ABTree{
					&ABTree{
						A:      4,
						B:      8,
						degree: 5,
						height: 0,
						keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7),
							IntKey(9)},
						values: []StoredObject{1, 4, 6, 7, 9},
						next:   []*ABTree{nil, nil, nil, nil, nil},
					},
					&ABTree{
						A:      4,
						B:      8,
						degree: 4,
						height: 0,
						keys:   []SearchKey{IntKey(10), IntKey(12), IntKey(13), IntKey(14)},
						values: []StoredObject{10, 12, 13, 14},
						next:   []*ABTree{nil, nil, nil, nil},
					}},
				values: []StoredObject{nil, nil},
			},
			deleteKey:    IntKey(10),
			expectObj:    10,
			expectStatus: DeleteOk,
			expectTree: &ABTree{
				A:      4,
				B:      8,
				degree: 2,
				height: 1,
				keys:   []SearchKey{nil, IntKey(9)},
				next: []*ABTree{
					&ABTree{
						A:      4,
						B:      8,
						degree: 4,
						height: 0,
						keys:   []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7)},
						values: []StoredObject{1, 4, 6, 7},
						next:   []*ABTree{nil, nil, nil, nil},
					},
					&ABTree{
						A:      4,
						B:      8,
						degree: 4,
						height: 0,
						keys:   []SearchKey{IntKey(9), IntKey(12), IntKey(13), IntKey(14)},
						values: []StoredObject{9, 12, 13, 14},
						next:   []*ABTree{nil, nil, nil, nil},
					}},
				values: []StoredObject{nil, nil},
			},
		},
		{
			inputTree: &ABTree{
				A:      4,
				B:      8,
				degree: 2,
				height: 1,
				keys:   []SearchKey{nil, IntKey(9)},
				next: []*ABTree{
					&ABTree{
						A:      4,
						B:      8,
						degree: 4,
						height: 0,
						keys:   []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7)},
						values: []StoredObject{1, 4, 6, 7},
						next:   []*ABTree{nil, nil, nil, nil},
					},
					&ABTree{
						A:      4,
						B:      8,
						degree: 4,
						height: 0,
						keys:   []SearchKey{IntKey(9), IntKey(12), IntKey(13), IntKey(14)},
						values: []StoredObject{9, 12, 13, 14},
						next:   []*ABTree{nil, nil, nil, nil},
					}},
				values: []StoredObject{nil, nil},
			},
			deleteKey:    IntKey(9),
			expectObj:    9,
			expectStatus: DeleteOk,
			expectTree: &ABTree{
				A:      4,
				B:      8,
				degree: 7,
				height: 0,
				keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7), IntKey(12),
					IntKey(13), IntKey(14)},
				values: []StoredObject{1, 4, 6, 7, 12, 13, 14},
				next:   []*ABTree{nil, nil, nil, nil, nil, nil, nil},
			},
		},
	}

	for _, tt := range tests {
		tree := tt.inputTree
		obj, status := tree.Delete(tt.deleteKey)
		if status != tt.expectStatus {
			t.Fatalf("unexpected delete status: got=%t, want=%t",
				status, tt.expectStatus)
		}
		if !reflect.DeepEqual(obj, tt.expectObj) {
			t.Fatalf("unexpected del object returned: got=%+v, want=%+v",
				obj, tt.expectObj)
		}
		if !reflect.DeepEqual(tree, tt.expectTree) {
			t.Fatalf("unexpected tree state:\ngot= %s\nwant=%s",
				tree, tt.expectTree)
		}
	}
}

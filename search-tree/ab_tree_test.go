package search_tree

import (
	"reflect"
	"testing"
)

func TestABTreeFind(t *testing.T) {
	tree := &ABTree{
		a:      4,
		b:      8,
		degree: 4,
		height: 1,
		keys:   []SearchKey{nil, IntKey(10), IntKey(20), IntKey(50)},
		next: []*ABTree{
			&ABTree{
				a:      4,
				b:      8,
				degree: 5,
				height: 0,
				keys: []SearchKey{IntKey(1), IntKey(4), IntKey(6), IntKey(7),
					IntKey(9)},
				values: []StoredObject{1, 4, 6, 7, 9},
			},
			&ABTree{
				a:      4,
				b:      8,
				degree: 7,
				height: 0,
				keys: []SearchKey{IntKey(10), IntKey(12), IntKey(13), IntKey(14),
					IntKey(15), IntKey(17), IntKey(19)},
				values: []StoredObject{10, 12, 13, 14, 15, 17, 19},
			},
			&ABTree{
				a:      4,
				b:      8,
				degree: 4,
				height: 0,
				keys:   []SearchKey{IntKey(21), IntKey(23), IntKey(24), IntKey(42)},
				values: []StoredObject{21, 23, 24, 42},
			},
			&ABTree{
				a:      4,
				b:      8,
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

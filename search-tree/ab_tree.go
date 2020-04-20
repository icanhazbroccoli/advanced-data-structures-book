package search_tree

type ABTree struct {
	keys   []SearchKey
	values []StoredObject
	next   []*ABTree
	a, b   int
	degree int
	height int
}

var _ SearchTree = (*ABTree)(nil)

func NewABTree(a, b int) *ABTree {
	return &ABTree{
		a:      a,
		b:      b,
		keys:   make([]SearchKey, b),
		values: make([]StoredObject, b),
		next:   make([]*ABTree, b),
	}
}

func (t *ABTree) Find(key SearchKey) (StoredObject, FindStatus) {
	cur := t
	for cur.height >= 0 {
		lower := 0
		upper := cur.degree
		for upper > lower+1 {
			if key.LessThan(cur.keys[(upper+lower)/2]) {
				upper = (upper + lower) / 2
			} else {
				lower = (upper + lower) / 2
			}
		}
		if cur.height > 0 {
			cur = cur.next[lower]
		} else {
			if cur.keys[lower].EqualsTo(key) {
				return cur.values[lower], FindOk
			}
			break
		}
	}
	return nil, FindNone
}

func (t *ABTree) Insert(key SearchKey, value StoredObject) InsertStatus {
	return InsertNone
}

func (t *ABTree) Delete(key SearchKey) (StoredObject, DeleteStatus) {
	return nil, DeleteNone
}

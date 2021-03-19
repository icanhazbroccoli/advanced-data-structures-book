package hashmap

type node struct {
	key   Key
	value Value
	next  *node
}

type HashMap struct {
	size     int
	table    []*node
	hashfunc func(int, Key) int
}

func NewHashMap(size int) *HashMap {
	return &HashMap{
		size:  size,
		table: make([]*node, size),
	}
}

func (hm *HashMap) Find(key Key) (Value, bool) {
	ix := hm.hashfunc(hm.size, key)
	ptr := hm.table[ix]
	for ptr != nil {
		if ptr.key == key {
			return ptr.value, true
		}
		ptr = ptr.next
	}
	return nil, false
}

func (hm *HashMap) Insert(key Key, value Value) {
	ix := hm.hashfunc(hm.size, key)
	ptr := hm.table[ix]
	for ptr != nil {
		if ptr.key == key {
			ptr.value = value
			return
		}
	}
	newnode := &node{
		key:   key,
		value: value,
		next:  hm.table[ix],
	}
	hm.table[ix] = newnode
}

func (hm *HashMap) Delete(key Key) {
	ix := hm.hashfunc(hm.size, key)
	var prev *node
	ptr := hm.table[ix]
	for ptr != nil {
		if ptr.key == key {
			if prev != nil {
				prev.next = ptr.next
			} else {
				hm.table[ix] = ptr.next
			}
			return
		}
		prev = ptr
		ptr = ptr.next
	}
}

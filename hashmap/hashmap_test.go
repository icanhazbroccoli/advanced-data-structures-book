package hashmap

import "testing"

func TestHashmap(t *testing.T) {
	hm := NewHashMap(16)
	hm.hashfunc = func(size int, key Key) int {
		sum := 0
		for _, ch := range key {
			sum += int(ch)
		}
		return sum % size
	}
	hm.Insert("foo", 1)
	hm.Insert("bar", 2)
	hm.Insert("boo", 42)
	hm.Delete("bar")
	hm.Delete("non existing")
	hm.Find("foo")
}

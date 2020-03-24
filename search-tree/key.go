package search_tree

import "fmt"

type IntKey int64

var _ Comparable = (IntKey)(0)

func (i IntKey) LessThan(c Comparable) bool {
	i2, ok := c.(IntKey)
	if !ok {
		panic(fmt.Sprintf("comparison keys are of different types: %T vs %T",
			i, c))
	}
	return i < i2
}

func (i IntKey) EqualsTo(c Comparable) bool {
	i2, ok := c.(IntKey)
	if !ok {
		panic(fmt.Sprintf("comparison keys are of different types: %T vs %T",
			i, c))
	}
	return i == i2
}

func (i IntKey) LessThanOrEqualsTo(c Comparable) bool {
	return i.EqualsTo(c) || i.LessThan(c)
}

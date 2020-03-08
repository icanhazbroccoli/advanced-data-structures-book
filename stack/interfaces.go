package stack

type Item interface{}

type Stack interface {
	Empty() bool
	Push(Item)
	Pop() Item
	Peek() Item
}

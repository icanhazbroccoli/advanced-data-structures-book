package stack

import "errors"

type Item interface{}

type Stack interface {
	Empty() bool
	Push(Item) error
	Pop() (Item, error)
	Peek() (Item, error)
}

var (
	StackIsFullError = errors.New(
		"stack has achieved it's max capacity")
	StackIsEmptyError = errors.New(
		"stack is empty")
)

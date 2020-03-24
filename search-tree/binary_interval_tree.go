package search_tree

type BinaryIntervalTree struct {
	*BinaryTree
}

var _ IntervalSearchTree = (*BinaryIntervalTree)(nil)

func NewBinaryIntervalTree() *BinaryIntervalTree {
	return &BinaryIntervalTree{&BinaryTree{}}
}

func (t *BinaryIntervalTree) FindInterval(a, b SearchKey) []StoredObject {
	res := []StoredObject{}
	// conventionally a <= b
	if b.LessThan(a) {
		return res
	}

	stack := []*BinaryTree{}
	stack = append(stack, t.BinaryTree)
	var node *BinaryTree
	for len(stack) > 0 {
		node, stack = stack[0], stack[1:]
		if node.isLeaf() {
			if a.LessThanOrEqualsTo(node.key) && node.key.LessThan(b) {
				res = append(res, node.value)
			}
		} else if b.LessThanOrEqualsTo(node.key) {
			stack = append(stack, node.left)
		} else if node.key.LessThanOrEqualsTo(a) {
			stack = append(stack, node.right)
		} else {
			stack = append(stack, node.left, node.right)
		}
	}

	return res
}

package search_tree

// The argument is a linked list of SearchNode items with .right pointing to the
// next item in the list.
func MakeTreeBottomUp(list *BinaryTree) *BinaryTree {
	if list == nil {
		return NewBinaryTree()
	} else if list.right == nil {
		return list
	}
	end := NewBinaryTree()
	root := end
	end.left = list
	end.key = list.key
	list = list.right
	end.left.right = nil
	for list != nil {
		end.right = NewBinaryTree()
		end = end.right
		end.left = list
		end.key = list.key
		list = list.right
		end.left.right = nil
	}
	end.right = nil

	oldList := root
	var tmp1, tmp2, newList *BinaryTree
	for oldList.right != nil {
		tmp1 = oldList
		tmp2 = oldList.right
		oldList = oldList.right.right
		tmp2.right = tmp2.left
		tmp2.left = tmp1.left
		tmp1.left = tmp2
		tmp1.right = nil
		newList = tmp1
		end = tmp1
		for oldList != nil {
			if oldList.right == nil {
				end.right = oldList
				oldList = nil
			} else {
				tmp1 = oldList
				tmp2 = oldList.right
				oldList = oldList.right.right
				tmp2.right = tmp2.left
				tmp2.left = tmp1.left
				tmp1.left = tmp2
				tmp1.right = nil
				end.right = tmp1
				end = end.right
			}
		}
		oldList = newList
	}
	root = oldList.left

	return root
}

func MakeTreeTopDown(list *BinaryTree) *BinaryTree {
	if list == nil {
		return NewBinaryTree()
	}
	type stackItem struct {
		node1, node2 *BinaryTree
		number       int
	}
	length := 0
	for tmp := list; tmp != nil; tmp = tmp.right {
		length++
	}
	root := NewBinaryTree()
	var left, right stackItem
	current := stackItem{
		node1:  root,
		number: length,
	}
	stack := []stackItem{current}

	for len(stack) > 0 {
		current, stack = stack[0], stack[1:]
		if current.number > 1 {
			left.node1 = NewBinaryTree()
			left.node2 = current.node2
			left.number = current.number / 2
			right.node1 = NewBinaryTree()
			right.node2 = current.node1
			right.number = current.number - left.number
			current.node1.left = left.node1
			current.node1.right = right.node1
			stack = append([]stackItem{left, right}, stack...)
		} else {
			current.node1.value = list.value
			current.node1.key = list.key
			current.node1.left = nil
			current.node1.right = nil
			if current.node2 != nil {
				current.node2.key = list.key
			}
			list = list.right
		}
	}
	return root
}

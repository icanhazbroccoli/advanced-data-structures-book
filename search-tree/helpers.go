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

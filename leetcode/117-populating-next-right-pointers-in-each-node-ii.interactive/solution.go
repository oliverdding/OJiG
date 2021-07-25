package leetcode

func findFirstChildInLevel(root *Node) *Node {
	for root != nil {
		if root.Left != nil {
			return root.Left
		}
		if root.Right != nil {
			return root.Right
		}
		root = root.Next
	}
	return nil
}

func connect(root *Node) *Node {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	// connect own child
	var child *Node = nil
	if root.Left != nil {
		root.Left.Next = root.Right
		child = root.Left
	}
	if root.Right != nil {
		child = root.Right
	}
	// Connect child with next that has child
	if child != nil {
		cur := root.Next
		for cur != nil {
			if cur.Left != nil {
				child.Next = cur.Left
				break
			}
			if cur.Right != nil {
				child.Next = cur.Right
				break
			}
			cur = cur.Next
		}
	}
	// interactive
	connect(root.Right)
	connect(root.Left)
	return root
}

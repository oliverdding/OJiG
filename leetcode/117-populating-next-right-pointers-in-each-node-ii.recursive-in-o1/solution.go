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
	if root == nil {
		return nil
	}
	cur := root
	nextLineHead := findFirstChildInLevel(root)
	for cur != nil {
		var child *Node
		// Connect own child
		if cur.Left != nil {
			cur.Left.Next = cur.Right
			child = cur.Left
		}
		if cur.Right != nil {
			child = cur.Right
		}
		// Connect child with first child in level
		if cur.Next == nil { // last node in this level
			cur = nextLineHead
			nextLineHead = findFirstChildInLevel(cur)
			continue
		}
		if child != nil {
			child.Next = findFirstChildInLevel(cur.Next)
		}
		cur = cur.Next
	}
	return root
}

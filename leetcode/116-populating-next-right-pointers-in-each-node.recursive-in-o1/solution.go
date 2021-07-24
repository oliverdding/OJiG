package leetcode

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	cur := root
	nextLineHead := cur.Left
	for cur.Left != nil {
		cur.Left.Next = cur.Right
		if cur.Next == nil { // last node in this level
			cur = nextLineHead
			nextLineHead = cur.Left
			continue
		}
		cur.Right.Next = cur.Next.Left
		cur = cur.Next
	}
	return root
}

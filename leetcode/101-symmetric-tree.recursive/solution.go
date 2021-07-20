package leetcode

func isSymmetric(root *TreeNode) bool {
	var (
		next  []*TreeNode
		left  *TreeNode
		right *TreeNode
	)
	next = append(next, root, root)
	for len(next) != 0 {
		left = next[len(next)-1]
		right = next[len(next)-2]
		next = next[:len(next)-2]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		next = append(next, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}

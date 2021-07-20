package leetcode

/*
  看了题解：递归
  ---
  自己想太复杂了，用非递归方式先序遍历左子树，后续遍历右子树。
*/
func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

func check(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
}

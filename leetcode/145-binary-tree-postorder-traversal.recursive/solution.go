package leetcode

/*
  递归版后序序遍历
*/
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return append(append(append([]int{}, postorderTraversal(root.Left)...), postorderTraversal(root.Right)...), root.Val)
}

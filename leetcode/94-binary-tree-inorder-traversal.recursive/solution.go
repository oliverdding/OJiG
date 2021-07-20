package leetcode

/*
  递归版中序遍历
*/
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return append(inorderTraversal(root.Left), append([]int{root.Val}, inorderTraversal(root.Right)...)...)
}

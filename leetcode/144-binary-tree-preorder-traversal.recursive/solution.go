package leetcode

/*
  递归版前序遍历
*/
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return append([]int{root.Val}, append(preorderTraversal(root.Left), preorderTraversal(root.Right)...)...)
}

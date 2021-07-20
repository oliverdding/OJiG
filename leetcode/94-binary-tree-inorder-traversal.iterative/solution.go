package leetcode

/*
  迭代版中序遍历
  对于每一个节点，不断循环
  左子树访问完成 -> 本节点访问完成 -> 访问右子树
*/
func inorderTraversal(root *TreeNode) (result []int) {
	var (
		stack []*TreeNode
	)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		root = root.Right
	}
	return
}

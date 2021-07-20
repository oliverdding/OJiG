package leetcode

/*
  迭代版前序遍历
  对于每一个节点，不断循环
  本节点访问完成 -> 左子树访问完成 -> 访问右子树
*/
func preorderTraversal(root *TreeNode) (result []int) {
	var (
		stack []*TreeNode
	)
	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}

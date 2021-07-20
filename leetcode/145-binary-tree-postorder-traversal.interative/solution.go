package leetcode

/*
  迭代版后序遍历
  对于每一个节点，不断循环
  左子树访问完成 -> 右子树访问完成 -> 访问本节点
*/
func postorderTraversal(root *TreeNode) (result []int) {
	var (
		stack []*TreeNode
		prev  *TreeNode // 用于记录右子树是否被访问过
	)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev { // 可以访问val：1. 右子树为空；2. 右子树被访问过；
			result = append(result, root.Val)
			prev = root
			root = nil
		} else { // 优先访问右子树：切换环境到右子树，压栈处理
			stack = append(stack, root)
			root = root.Right
		}
	}
	return
}

package leetcode

import (
	"math"
)

var base int64

/*
  中序遍历中，首先遍历左子树，然后中间节点，最后右子数。
  对符合二叉搜索树定义的数结构进行中序遍历肯定可以得到一个正序数组。
*/
func isValidBST(root *TreeNode) bool {
	base = math.MinInt64
	return inorderTraversal(root)
}

func inorderTraversal(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if inorderTraversal(root.Left) != true {
		return false
	}
	if int64(root.Val) <= base {
		return false
	}
	base = int64(root.Val)
	if inorderTraversal(root.Right) != true {
		return false
	}
	return true
}

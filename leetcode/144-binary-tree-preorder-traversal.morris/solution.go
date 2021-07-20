package leetcode

/*
o(n)时间，o(1)空间的 Morris 前序遍历:

需要两个指针：cur、rightMost
1. cur无左子树：访问cur，cur向右移动
2. cur有左子树：找到cur左子树最右叶子节点
    1. 如果rightMost的右指针指向空，访问cur，令指向cur，cur向左移动
    2. 如果rightMost的右指针指向cur，令指向空，cur向右移动

总结一下：思想是通过每个节点的左子树最右叶子节点的右指针指回本节点，以此方式得以返回之前的节点，得以进入右子树
*/

func preorderTraversal(root *TreeNode) (vals []int) {
	var cur, rightMost *TreeNode = root, nil
	for cur != nil {
		if rightMost = cur.Left; rightMost == nil { // 1
			vals = append(vals, cur.Val)
		} else {
			for rightMost.Right != nil && rightMost.Right != cur { // 2
				rightMost = rightMost.Right
			}
			if rightMost.Right == nil { // 2.1
				vals = append(vals, cur.Val)
				rightMost.Right = cur
				cur = cur.Left
				continue
			}
			// 2.2
			rightMost.Right = nil
		}
		cur = cur.Right
	}
	return
}

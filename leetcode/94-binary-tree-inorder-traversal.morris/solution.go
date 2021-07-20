package leetcode

/*
o(n)时间，o(1)空间的 Morris 中序遍历:

需要两个指针：cur、rightMost
1. cur无左子树：访问cur，cur向右移动
2. cur有左子树：找到cur左子树最右叶子节点
    1. 如果rightMost的右指针指向空，令指向cur，cur向左移动
    2. 如果rightMost的右指针指向cur，访问cur，令指向空，cur向右移动

总结一下：与前序遍历相比，仅仅改变访问中间节点时机即可。
*/

func inorderTraversal(root *TreeNode) (vals []int) {
	var cur, rightMost *TreeNode = root, nil
	for cur != nil {
		if rightMost = cur.Left; rightMost == nil {
			vals = append(vals, cur.Val)
		} else {
			for rightMost.Right != nil && rightMost.Right != cur {
				rightMost = rightMost.Right
			}
			if rightMost.Right == nil {
				rightMost.Right = cur
				cur = cur.Left
				continue
			}
			vals = append(vals, cur.Val)
			rightMost.Right = nil
		}
		cur = cur.Right
	}
	return
}

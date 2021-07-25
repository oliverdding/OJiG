package leetcode

func levelOrder(root *TreeNode) (vals [][]int) {
	if root == nil {
		return
	}
	nexLevel := &[]*TreeNode{root}
	for len(*nexLevel) != 0 {
		var (
			level []int
		)
		curLevel := nexLevel
		nexLevel = new([]*TreeNode)
		for i := 0; i < len(*curLevel); i++ {
			level = append(level, (*curLevel)[i].Val)
			if (*curLevel)[i].Left != nil {
				*nexLevel = append(*nexLevel, (*curLevel)[i].Left)
			}
			if (*curLevel)[i].Right != nil {
				*nexLevel = append(*nexLevel, (*curLevel)[i].Right)
			}
		}
		vals = append(vals, level)
	}
	return
}

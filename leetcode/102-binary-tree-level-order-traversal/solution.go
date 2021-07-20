package leetcode

func levelOrder2(root *TreeNode) (vals [][]int) {
	if root == nil {
		return
	}
	curLevel := new([]*TreeNode)
	nexLevel := new([]*TreeNode)
	*curLevel = append(*curLevel, root)
	for len(*curLevel) != 0 {
		var (
			level []int
		)
		for i := 0; i < len(*curLevel); i++ {
			level = append(level, (*curLevel)[i].Val)
			if (*curLevel)[i].Left != nil {
				*nexLevel = append(*nexLevel, (*curLevel)[i].Left)
			}
			if (*curLevel)[i].Right != nil {
				*nexLevel = append(*nexLevel, (*curLevel)[i].Right)
			}
		}
		*curLevel = (*curLevel)[:0]
		vals = append(vals, level)
		curLevel, nexLevel = nexLevel, curLevel
	}
	return
}

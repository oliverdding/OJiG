package leetcode

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	curLevel := new([]*Node)
	nexLevel := new([]*Node)
	*curLevel = append(*curLevel, root)
	for len(*curLevel) != 0 {
		last := len(*curLevel) - 1
		for i := 0; i < last; i++ {
			(*curLevel)[i].Next = (*curLevel)[i+1]
			if (*curLevel)[i].Left != nil {
				*nexLevel = append(*nexLevel, (*curLevel)[i].Left)
			}
			if (*curLevel)[i].Right != nil {
				*nexLevel = append(*nexLevel, (*curLevel)[i].Right)
			}
		}
		if (*curLevel)[last].Left != nil {
			*nexLevel = append(*nexLevel, (*curLevel)[last].Left)
		}
		if (*curLevel)[last].Right != nil {
			*nexLevel = append(*nexLevel, (*curLevel)[last].Right)
		}
		*curLevel = (*curLevel)[:0]
		curLevel, nexLevel = nexLevel, curLevel
	}
	return root
}

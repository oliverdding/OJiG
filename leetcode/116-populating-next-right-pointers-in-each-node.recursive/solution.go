package leetcode

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	nexLevel := &[]*Node{root}
	for len(*nexLevel) != 0 {
		curLevel := nexLevel
		nexLevel = new([]*Node)
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
	}
	return root
}

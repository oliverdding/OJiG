package leetcode

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	val := postorder[len(postorder)-1]
	leftSize := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == val {
			leftSize = i
			break
		}
	}
	return &TreeNode{
		Val:   val,
		Left:  buildTree(inorder[:leftSize], postorder[:leftSize]),
		Right: buildTree(inorder[leftSize+1:], postorder[leftSize:len(postorder)-1]),
	}
}

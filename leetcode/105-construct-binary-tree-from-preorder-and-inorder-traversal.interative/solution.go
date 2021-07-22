package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	val := preorder[0]
	leftSize := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == val {
			leftSize = i
			break
		}
	}
	return &TreeNode{
		Val:   val,
		Left:  buildTree(preorder[1:leftSize+1], inorder[:leftSize]),
		Right: buildTree(preorder[leftSize+1:], inorder[leftSize+1:]),
	}
}

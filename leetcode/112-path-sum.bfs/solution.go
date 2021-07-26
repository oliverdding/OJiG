package leetcode

// Breadth First Search
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var vec = []*TreeNode{root}
	var sum = []int{root.Val}
	for len(vec) != 0 {
		size := len(vec)
		for i := 0; i < size; i++ {
			node := vec[i]
			val := sum[i]
			if node.Left == nil && node.Right == nil {
				if val == targetSum {
					return true
				}
				continue
			}
			if node.Left != nil {
				vec = append(vec, node.Left)
				sum = append(sum, val+node.Left.Val)
			}
			if node.Right != nil {
				vec = append(vec, node.Right)
				sum = append(sum, val+node.Right.Val)
			}
		}
		vec = vec[size:]
		sum = sum[size:]
	}
	return false
}

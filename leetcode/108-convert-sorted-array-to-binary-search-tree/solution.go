package leetcode

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   nums[n/2],
		Left:  nil,
		Right: nil,
	}
	if n != 1 {
		root.Left = sortedArrayToBST(nums[0 : n/2])
		root.Right = sortedArrayToBST(nums[n/2+1:])
	}
	return root
}

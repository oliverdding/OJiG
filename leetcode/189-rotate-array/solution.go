package leetcode

/*
  仔细观察可以看出，循环右移的结果等同于反转后部分反转
  [1,2,3,4] -> 2
  1. 反转 [4,3,2,1]
  2. 部分反转 [3,4,1,2]
*/
func rotate(nums []int, k int) {
	length := len(nums)
	k = k % length
	revert(nums[:])
	revert(nums[:k])
	revert(nums[k:])
}

func revert(nums []int) {
	for b, e := 0, len(nums)-1; b <= e; b, e = b+1, e-1 {
		nums[b], nums[e] = nums[e], nums[b]
	}
}

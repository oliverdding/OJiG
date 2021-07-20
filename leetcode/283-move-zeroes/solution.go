package leetcode

/*
  想到双指针方式，1号指针一直走直到元素为0，此时2号指针找非零元素并拷贝到1号指针位置。
  重复。
*/
func moveZeroes(nums []int) {
	j := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			continue
		}
		for j = max(i+1, j); j < length; j++ {
			if nums[j] != 0 {
				nums[i], nums[j] = nums[j], 0
				break
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

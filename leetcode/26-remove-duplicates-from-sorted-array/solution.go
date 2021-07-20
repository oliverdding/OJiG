package leetcode

func removeDuplicates(nums []int) int {
	i := 0
	s := len(nums)
	for j := 1; j < s; j++ {
		if nums[i] != nums[j] {
			i = i + 1
			nums[i] = nums[j]
		}
	}
	return i + 1
}

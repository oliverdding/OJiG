package leetcode

/*
  还可以用用空间换时间。使用map数据结构，将以往的数值记录下来。发现重复即返回。
*/
func containsDuplicate(nums []int) bool {
	exist := make(map[int]bool)
	for _, num := range nums {
		if exist[num] {
			return true
		}
		exist[num] = true
	}
	return false
}

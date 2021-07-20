package leetcode

/*
  用map，needed->position
*/
func twoSum(nums []int, target int) []int {
	neededs := make(map[int]int)
	for i, num := range nums {
		if needed, ok := neededs[num]; ok {
			return []int{needed, i}
		}
		neededs[target-num] = i
	}
	return nil
}

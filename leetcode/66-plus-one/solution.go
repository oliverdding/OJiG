package leetcode

/*
  唯一需要注意的是进位，因此要从数组尾部开始处理
*/
func plusOne(digits []int) []int {
	var result []int
	tap := 1
	for i := len(digits) - 1; i >= 0; i-- {
		tap, digits[i] = (tap+digits[i])/10, (tap+digits[i])%10
	}
	if tap == 1 {
		result = make([]int, len(digits)+1)
		result[0] = 1
		copy(result[1:], digits)
	} else {
		result = make([]int, len(digits))
		copy(result, digits)
	}
	return result
}

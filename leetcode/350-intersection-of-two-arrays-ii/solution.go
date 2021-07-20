package leetcode

import (
	"sort"
)

/*
  先排序，两个排好序的数组更方便。
  1. 两个数组，两个指针。
  2. 值更小的指针前进。
  3. 遇到相同的值，取出，并两个指针同时前进
*/
func intersect(nums1 []int, nums2 []int) []int {
	var result []int
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			result = append(result, nums1[i])
			i++
			j++
		}
	}
	return result
}

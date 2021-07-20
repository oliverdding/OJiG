package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		i  = m + n
		i1 = m
		i2 = n
	)
	for i1 != 0 || i2 != 0 {
		if i1 == 0 {
			copy(nums1[:i], nums2[:i2])
			break
		}
		if i2 == 0 {
			break
		}
		if nums1[i1-1] <= nums2[i2-1] {
			nums1[i-1] = nums2[i2-1]
			i2--
		} else {
			nums1[i-1] = nums1[i1-1]
			i1--
		}
		i--
	}
}

package leetcode

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return            true if current version is bad
 *                    false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool {
	// not implementation
	return true
}

func firstBadVersion(n int) int {
	var (
		begin = 1
		end   = n
		mid   int
	)
	for {
		if begin == end {
			break
		}
		mid = (begin + end) / 2
		if isBadVersion(mid) {
			end = mid
		} else {
			begin = mid + 1
		}
	}
	return begin
}

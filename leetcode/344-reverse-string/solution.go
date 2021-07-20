package leetcode

func reverseString(s []byte) {
	n := len(s)
	for i := n/2 - 1; i >= 0; i-- {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}

package leetcode

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	i := 0
	for i < len(strs[0]) {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				goto RET
			}
		}
		i++
	}
RET:
	return strs[0][:i]
}

package leetcode

/*
  暴力算法最大的问题是会重复搜索。
  因此可以使用map记录下每个字符出现的次数。优化可以考虑把字符数组当作普通数值数组排序后寻找。
  又由于提示给出，只存储小写字母，那么就可以使用26位数组当作map使用。
*/
func firstUniqChar(s string) int {
    arr := []byte(s)
    counts := [26]int{}
    for _, a := range arr {
        counts[a-'a']++
    }
    for i, a := range arr {
        if counts[a-'a'] == 1 {
            return i
        }
    }
    return -1
}

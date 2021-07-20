package leetcode

/*
  大名鼎鼎的KMP算法
*/
func strStr2(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    } else if len(haystack) == 0 {
        return -1
    }
    preLength := calPreLength(needle)
    hlen, nlen := len(haystack), len(needle)
    for i, k := 0, 0; i < hlen; i++ {
        for k != 0 && haystack[i] != needle[k] {
            k = preLength[k-1]
        }
        if haystack[i] == needle[k] {
            k++
        }
        if k == nlen {
            return i + 1 - nlen
        }
    }
    return -1
}

func calPreLength(s string) []int {
    n := len(s)
    preLength := make([]int, n)
    k := 0
    for i := 1; i < n; i++ {
        for k != 0 && s[k] != s[i] {
            k = preLength[k-1]
        }
        if s[k] == s[i] {
            k++
        }
        preLength[i] = k
    }
    return preLength
}

package leetcode

/*
  也就是统计两个字符串中字符种类数量是否相等，可以用map记录。
  由于条件给出，都是小写字母，所以用数组来记录
*/
func isAnagram(s string, t string) bool {
    var counts [26]int
    if len(s) != len(t) {
        return false
    }
    for i := range s {
        counts[s[i]-'a']++
        counts[t[i]-'a']--
    }
    for _, count := range counts {
        if count != 0 {
            return false
        }
    }
    return true
}

package leetcode

/*
  只考虑字母和数字字符，忽略大小写，判断回文。
  首先判断回文，我们可以用双指针，可以反转，可以用栈。
  用栈解决可以先将相关的字符取出。
  这里使用双指针解决。
*/
func isPalindrome(s string) bool {
    if len(s) == 0 {
        return true
    }
    bp, ep := 0, len(s)-1
    for bp <= ep {
        // 这里用continue而不是for，为了讲过程慢下来
        if !isValidate(s[bp]) {
            bp++
            continue
        }
        if !isValidate(s[ep]) {
            ep--
            continue
        }
        if !isEqual(s[bp], s[ep]) {
            return false
        }
        bp++
        ep--
    }
    return true
}

func isEqual(a, b uint8) bool {
    if a <= 'Z' && a >= 'A' {
        a += 32
    }
    if b <= 'Z' && b >= 'A' {
        b += 32
    }
    return a == b
}

func isValidate(ch byte) bool {
    return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

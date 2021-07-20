package leetcode

import (
    "math"
)

/*
  想到自动机：

  1 等待空格
  2 等待数字

  1状态，遇到空格->1
  1状态，遇到'-'、'+'->2
  1状态，遇到数字->2
  1状态，遇到其他->return 0
  2状态，遇到数字->2
  2状态，遇到其他->return

  由于状态很少，并且不存在2转1的路径，直接扫描
*/
func myAtoi(s string) int {
    var (
        isNeg = false
        slen  = len(s)
        i     = 0
    )
    
    // 吃掉空格
    for ; i < slen; i++ {
        if s[i] != ' ' {
            break
        }
    }
    if i == slen {
        return 0
    }
    
    // 吃掉可能存在的符号
    if s[i] == '-' {
        isNeg = true
        i++
    } else if s[i] == '+' {
        i++
    }
    
    // 吃掉数字
    result := int32(0) // OJ32位，系统64位
    for ; i < slen; i++ {
        if s[i] > '9' || s[i] < '0' {
            break
        }
        if result < math.MaxInt32/10 ||
            result == math.MaxInt32/10 && s[i]-'0' <= 7 {
            result = result*10 + int32(s[i]-'0')
            continue
        }
        if isNeg {
            return math.MinInt32
        } else {
            return math.MaxInt32
        }
    }
    
    if isNeg {
        result = -result
    }
    return int(result)
}

package leetcode

import (
    "strings"
)

func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    tmp := countAndSay(n - 1)
    var result strings.Builder
    count := 1
    // 解析tmp
    for i := 0; i < len(tmp); i++ {
        if i == len(tmp)-1 || tmp[i] != tmp[i+1] {
            result.WriteByte(byte(count + '0'))
            result.WriteByte(tmp[i])
            count = 1
        } else {
            count++
        }
    }
    return result.String()
}

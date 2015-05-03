package stringutil

import "strings"

// WordCount 获取一句话中每个英文单词出现的次数
func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    ret := make(map[string]int)
    for i:=0; i<len(words); i++ {
        w := words[i]
        v,o := ret[w]
        if !o {
            ret[w] = 1
        } else {
            ret[w] = v + 1
        }
    }
    return ret
}


// Fibonacci 函数会返回一个返回 int 的函数。
func Fibonacci() func() int {
    x := 0
    x1 := 1
    i := 0
    return func() int {
        if i <= 1 {
            i++
            return i-1
        }
        ret := x + x1
        x = x1
        x1 = ret
        return ret
    }
}

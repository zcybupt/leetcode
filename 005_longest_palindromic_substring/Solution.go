package main

import (
    "fmt"
    "math"
)

func preprocess(s string) string {
    result := "^"
    for _, ele := range s {
        result += "#" + string(ele)
    }

    return result + "#$"
}

func longestPalindrome(s string) string {
    if len(s) == 0 {
        return ""
    }

    tmpStr := preprocess(s)
    strLen := len(tmpStr)
    p := make([]int, strLen)
    C, R := 0, 0
    index, maxLen := 0, 0

    for i := 1; i < strLen-1; i++ {
        if R > i {
            p[i] = int(math.Min(float64(p[2*C-i]), float64(R-i)))
        } else {
            p[i] = 1
        }

        for tmpStr[i+p[i]] == tmpStr[i-p[i]] {
            p[i]++
        }

        if R < i+p[i] {
            R = i + p[i]
            C = i
        }

        if maxLen < p[i]-1 {
            maxLen = p[i] - 1
            index = i
        }
    }

    start := (index - maxLen) / 2
    return s[start: start+maxLen]
}

func main() {
    fmt.Println(longestPalindrome("babad"))
}

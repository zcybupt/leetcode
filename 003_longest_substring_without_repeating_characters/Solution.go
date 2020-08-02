package main

import (
    "fmt"
    "math"
)

func lengthOfLongestSubstring(s string) int {
    var window [128]int8
    length := len(s)
    rk, ans := -1, 0

    for i, _ := range s {
        if i != 0 {
            window[s[i-1]]--
        }
        for rk+1 < length && window[s[rk+1]] == 0 {
            window[s[rk+1]] = 1
            rk++
        }
        ans = int(math.Max(float64(ans), float64(rk-i+1)))
    }
    return ans
}

func main() {
    fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

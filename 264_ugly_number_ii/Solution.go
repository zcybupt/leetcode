package main

import "fmt"

func nthUglyNumber(n int) int {
    dp := []int{1}
    a, b, c := 0, 0, 0
    for i := 1; i < n; i++ {
        n2, n3, n5 := dp[a] * 2, dp[b] * 3, dp[c] * 5
        dp = append(dp, min(min(n2, n3), n5))
        if dp[i] == n2 { a++ }
        if dp[i] == n3 { b++ }
        if dp[i] == n5 { c++ }
    }

    return dp[n - 1]
}

func min(a, b int) int {
    if a < b { return a }
    return b
}

func main() {
    fmt.Println(nthUglyNumber(10))
}

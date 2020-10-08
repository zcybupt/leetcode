package main

import "fmt"

func lengthOfLIS(nums []int) int {
    arrLen := len(nums)
    if arrLen < 2 { return arrLen }

    dp := make([]int, arrLen)
    dp[0] = 1

    for i := 1; i < arrLen; i++ {
        dp[i] = 1
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        }
    }

    result := 0
    for i := 0; i < arrLen; i++ { result = max(result, dp[i]) }

    return result
}

func max(a, b int) int {
    if a > b { return a }
    return b
}

func main() {
    nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
    fmt.Println(lengthOfLIS(nums))
}

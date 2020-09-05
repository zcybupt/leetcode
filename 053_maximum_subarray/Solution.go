package main

import "fmt"

func maxSubArray(nums []int) int {
    if nums == nil || len(nums) == 0 {
        return 0
    }

    current := 0
    maxSum := ^int(^uint(0) >> 1)

    for _, num := range nums {
        current += num
        maxSum = max(current, maxSum)
        current = max(current, 0)
    }

    return maxSum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
    fmt.Println(maxSubArray(nums))
}

package main

import "fmt"

func canJump(nums []int) bool {
    arrLen := len(nums)
    rightMost := 0

    for i := 0; i < arrLen; i++ {
        if i <= rightMost {
            rightMost = max(rightMost, i + nums[i])
            if rightMost >= arrLen - 1 {
                return true
            }
        }
    }

    return false
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    nums := []int{2, 3, 1, 1, 4}
    fmt.Println(canJump(nums))
}


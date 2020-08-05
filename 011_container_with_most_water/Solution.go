package main

import (
    "fmt"
)

func maxArea(height []int) int {
    if height == nil || len(height) == 0 {
        return 0
    }

    l, r := 0, len(height) - 1
    area, ans := 0, 0

    for l < r {
        if (height[l] < height[r]) {
            area = height[l] * (r - l)
            l++
        } else {
            area = height[r] * (r - l)
            r--
        }
        if ans < area {
            ans = area
        }
    }

    return ans
}

func main() {
    height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
    fmt.Println(maxArea(height))
}


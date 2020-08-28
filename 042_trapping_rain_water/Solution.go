package main

import "fmt"

func trap(heightArr []int) int {
    if heightArr == nil || len(heightArr) < 3 {
        return 0
    }

    arrLen := len(heightArr)
    leftMax, rightMax := heightArr[0], heightArr[arrLen - 1]
    L, R := 1, arrLen - 2
    result := 0

    for L <= R {
        if leftMax <= rightMax {
            result += max(0, leftMax - heightArr[L])
            leftMax = max(leftMax, heightArr[L])
            L++
        } else {
            result += max(0, rightMax - heightArr[R])
            rightMax = max(rightMax, heightArr[R])
            R--
        }
    }

    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    heightArr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
    fmt.Println(trap(heightArr))
}

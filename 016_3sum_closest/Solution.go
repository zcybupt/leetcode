package main

import (
    "fmt"
    "math"
    "sort"
)

func threeSumClosest(nums []int, target int) int {
    arrLen := len(nums)
    if nums == nil || arrLen < 3 {
        return 0
    }

    sort.Ints(nums)
    result, curSum, curDiff := 0, 0, 0
    minDiff := math.MaxInt32

    for i, v := range nums {
        L, R := i+1, arrLen-1
        for L < R {
            curSum = v + nums[L] + nums[R]
            if curSum == target {
                return curSum
            }
            curDiff = int(math.Abs(float64(curSum - target)))
            if curDiff < minDiff {
                result = curSum
                minDiff = curDiff
            }

            if curSum > target {
                R--
            } else {
                L++
            }
        }
    }

    return result
}

func main() {
    nums := []int{-1, 0, 1, 1, 55}
    target := 3
    fmt.Println(threeSumClosest(nums, target))
}


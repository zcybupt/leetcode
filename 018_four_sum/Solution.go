package main

import (
    "fmt"
    "sort"
)

func fourSum(nums []int, target int) [][]int {
    arrLen := len(nums)
    var results [][]int
    if arrLen < 4 {
        return results
    }

    sort.Ints(nums)
    if isExceeded(nums, target, 0, arrLen-1) {
        return results
    }

    for i := 0; i < arrLen - 3; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for j := i+1; j < arrLen-2; j++ {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            L, R := j+1, arrLen-1
            if isExceeded(nums, target, i, R) {
                continue
            }
            for L < R {
                tmpSum := nums[i] + nums[j] + nums[L] + nums[R]
                if tmpSum == target {
                    results = append(results, []int{nums[i], nums[j], nums[L], nums[R]})
                    for L < R && nums[L] == nums[L+1] {
                        L++
                    }
                    for L < R && nums[R] == nums[R-1] {
                        R--
                    }
                    L++
                    R--
                } else if tmpSum > target {
                    R--
                } else {
                    L++
                }
            }
        }
    }

    return results
}

func isExceeded(nums []int, target, L, R int) bool {
    minSum := nums[L] + nums[L+1] + nums[L+2] + nums[L+3]
    maxSum := nums[R] + nums[R-1] + nums[R-2] + nums[R-3]
    return target > maxSum || target < minSum
}

func main() {
    nums := []int{1, 0, -1, 0, -2, 2}
    target := 0
    fmt.Printf("%#v\n", fourSum(nums, target))
}

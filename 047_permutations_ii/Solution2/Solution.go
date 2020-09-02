package main

import (
    "fmt"
    "sort"
)

func permuteUnique(nums []int) [][]int {
    var results [][]int
    arrLen := len(nums)
    sort.Ints(nums)
    results = append(results, append([]int{}, nums...))

    for nextPermutation(nums, arrLen) {
        results = append(results, append([]int{}, nums...))
    }

    return results
}

func nextPermutation(nums []int, arrLen int) bool {
    i, j, k := arrLen - 2, arrLen - 1, arrLen - 1

    for i >= 0 && nums[i] >= nums[j] {
        i--
        j--
    }

    if i == -1 {
        return false
    }

    for nums[i] >= nums[k] {
        k--
    }
    nums[i], nums[k] = nums[k], nums[i]

    for i, j = j, arrLen - 1; i < j; i, j = i + 1, j - 1 {
        nums[i], nums[j] = nums[j], nums[i]
    }

    return true
}

func main() {
    nums := []int{3, 3, 0, 3}
    results := permuteUnique(nums)
    for _, result := range results {
        fmt.Println(result)
    }
}

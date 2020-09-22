package main

import "fmt"

func subsets(nums []int) [][]int {
    arrLen := len(nums)
    var results [][]int

    for i := 0; i <= arrLen; i++ {
        used := make([]bool, arrLen)
        dfs(nums, &results, 0, arrLen, i, used)
    }

    return results
}

func dfs(nums []int, results *[][]int, start, arrLen, steps int, used []bool) {
    if steps == 0 {
        var tmpRes []int
        for i := 0; i < arrLen; i++ {
            if used[i] {
                tmpRes = append(tmpRes, nums[i])
            }
        }
        *results = append(*results, tmpRes)
        return
    }

    for i := start; i < arrLen; i++ {
        used[i] = true
        dfs(nums, results, i + 1, arrLen, steps - 1, used)
        used[i] = false
    }
}

func main() {
    nums := []int{1, 2, 3}
    results := subsets(nums)
    for _, result := range results {
        fmt.Println(result)
    }
}

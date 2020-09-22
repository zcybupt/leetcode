package main

import "fmt"

func subsets(nums []int) [][]int {
    arrLen := len(nums)
    var results [][]int

    for i := 0; i <= arrLen; i++ {
        var path []int
        dfs(nums, 0, arrLen, i, &results, path)
    }

    return results
}

func dfs(nums []int, start, arrLen, steps int, results *[][]int, path []int) {
    if steps == 0 {
        *results = append(*results, append([]int{}, path...))
        return
    }

    for i := start; i < arrLen; i++ {
        path = append(path, nums[i])
        dfs(nums, i + 1, arrLen, steps - 1, results, path)
        path = path[:len(path) - 1]
    }
}

func main() {
    nums := []int{1, 2, 3}
    results := subsets(nums)
    for _, result := range results {
        fmt.Println(result)
    }
}

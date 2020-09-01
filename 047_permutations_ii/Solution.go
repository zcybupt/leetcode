package main

import (
    "fmt"
    "sort"
)

func permuteUnique(nums []int) [][]int {
    var results [][]int
    sort.Ints(nums)
    arrLen := len(nums)
    dfs(&results, nums, []int{}, make([]int, arrLen), arrLen)

    return results
}

func dfs(results *[][]int, nums, tmp, visited []int, arrLen int) {
    if len(tmp) == arrLen {
        *results = append(*results, append([]int{}, tmp...)) // 注意深拷贝
        return
    }

    for i := 0; i < arrLen; i++ {
        if visited[i] == 1 || (i > 0 && nums[i] == nums[i - 1] && visited[i - 1] != 1) {
            continue
        }
        tmp = append(tmp, nums[i])
        visited[i] = 1
        dfs(results, nums, tmp, visited, arrLen)
        tmp = tmp[:len(tmp) - 1]
        visited[i] = 0
    }
}

func main() {
    nums := []int{3, 3, 0, 3}
    results := permuteUnique(nums)
    for _, result := range results {
        fmt.Println(result)
    }
}


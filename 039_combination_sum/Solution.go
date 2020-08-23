package main

import (
    "fmt"
    "sort"
)

func combinationSum(candidates []int, target int) [][]int {
    arrLen := len(candidates)
    var results [][]int
    if arrLen == 0 {
        return results
    }

    var path []int
    sort.Ints(candidates)
    dfs(candidates, 0, arrLen, target, results, path)
    return results
}

func dfs(candidates []int, begin, arrLen, target int, results [][]int, path []int) {
    if target == 0 {
        results = append(results, path)
        return
    }

    for i := begin; i < arrLen; i++ {
        if target - candidates[i] < 0 {
            break
        }

        path = append(path, candidates[i])
        dfs(candidates, i, arrLen, target - candidates[i], results, path)
        path = path[:len(path) - 1]
    }
}

func main() {
    nums := []int{2, 3, 6, 7}
    target := 7
    results := combinationSum(nums, target)
    for _, result := range results {
        for _, ele := range result {
            fmt.Printf("%d ", ele)
        }
        fmt.Println();
    }
}

package main

import (
    "fmt"
    "sort"
)

var tmpNums []int
var arrLen int
var results [][]int

func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    tmpNums = nums
    arrLen = len(nums)
    results = results[0:0] // 全局变量LeetCode可能不会清空内存中的数据
    for i := 0; i <= arrLen; i++ {
        dfs(0, i, make([]bool, arrLen))
    }

    return results
}

func dfs(start, k int, used []bool) {
    if k == 0 {
        var tmpRes []int
        for i := 0; i < arrLen; i++ {
            if used[i] { tmpRes = append(tmpRes, tmpNums[i]) }
        }
        results = append(results, tmpRes)
    }

    for i := start; i < arrLen; i++ {
        if i > 0 && !used[i - 1] && tmpNums[i] == tmpNums[i - 1] { continue }
        used[i] = true
        dfs(i + 1, k - 1, used)
        used[i] = false
    }
}

func main() {
    nums := []int{4, 4, 1, 4}
    results := subsetsWithDup(nums)
    for _, result := range results {
        fmt.Println(result)
    }
}

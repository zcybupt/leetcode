package main

import "fmt"

func permute(nums []int) [][]int {
    var results [][]int
    dfs(&results, nums, []int{})

    return results
}

func dfs(results *[][]int, nums, tmp []int) {
    if len(nums) == 0 {
        *results = append(*results, tmp)
        return
    }

    for i := 0; i < len(nums); i++ {
        dfs(results, combinedArr(nums[:i], nums[i+1:]), combinedArr(tmp, []int{nums[i]}))
    }
}

func combinedArr(arr1, arr2 []int) []int {
    var result []int
    result = append(result, arr1...)
    result = append(result, arr2...)

    return result
}

func main() {
    nums := []int{1, 2, 3}
    results := permute(nums)
    for _, result := range results {
        fmt.Println(result)
    }
}


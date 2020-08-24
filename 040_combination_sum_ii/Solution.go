package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	arrLen := len(candidates)
	var results [][]int
	if arrLen == 0 {
		return results
	}

	var path []int
	sort.Ints(candidates)
	dfs(candidates, 0, arrLen, target, &results, path)
	return results
}

func dfs(candidates []int, begin, arrLen, target int, results *[][]int, path []int) {
	if target == 0 {
		*results = append(*results, append([]int{}, path...))
		return
	}

	for i := begin; i < arrLen; i++ {
		if target < candidates[i] {
			break
		}
		if i > begin && candidates[i] == candidates[i-1] {
			continue
		}

		path = append(path, candidates[i])
		dfs(candidates, i+1, arrLen, target-candidates[i], results, path)
		path = path[:len(path)-1]
	}
}

func main() {
	nums := []int{2, 5, 2, 1, 2}
	target := 5
	results := combinationSum2(nums, target)
	for _, result := range results {
		for _, ele := range result {
			fmt.Printf("%d ", ele)
		}
		fmt.Println()
	}
}


package main

import "fmt"

func combine(n int, k int) [][]int {
    var results [][]int
    var path []int

    dfs(&results, path, 1, n, k)

    return results
}

func dfs(results *[][]int, path []int, start, n, k int) {
    if k == 0 {
        *results = append(*results, append([]int{}, path...))
        return
    }

    for i := start; i <= n; i++ {
        path = append(path, i)
        dfs(results, path, i + 1, n, k - 1)
        path = path[:len(path) - 1]
    }
}

func main() {
    results := combine(4, 2)

    for _, result := range results {
        fmt.Println(result)
    }
}

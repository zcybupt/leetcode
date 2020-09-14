package main

import "fmt"

func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 {
                dp[0][0] = grid[0][0]
            } else if i == 0 {
                dp[0][j] = dp[0][j - 1] + grid[0][j]
            } else if j == 0 {
                dp[i][0] = dp[i - 1][0] + grid[i][0]
            } else {
                dp[i][j] = grid[i][j] + min(dp[i - 1][j], dp[i][j - 1])
            }
        }
    }

    return dp[m - 1][n - 1]
}

func min(a, b int) int {
    if a > b { return b }
    return a
}

func main() {
    grid := [][]int {
        {1, 3, 1},
        {1, 5, 1},
        {4, 2, 1},
    }

    fmt.Println(minPathSum(grid))
}

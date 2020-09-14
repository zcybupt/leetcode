package main

import "fmt"

func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 {
                continue 
            } else if i == 0 {
                grid[0][j] += grid[0][j - 1]
            } else if j == 0 {
                grid[i][0] += grid[i - 1][0]
            } else {
                grid[i][j] += min(grid[i - 1][j], grid[i][j - 1])
            }
        }
    }

    return grid[m - 1][n - 1]
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


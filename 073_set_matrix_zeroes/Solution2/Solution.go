package main

import "fmt"

func setZeroes(matrix [][]int)  {
    colFlag := false
    m, n := len(matrix), len(matrix[0])

    for i := 0; i < m; i++ {
        if matrix[i][0] == 0 { colFlag = true }
        for j := 1; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[0][j] = 0
                matrix[i][0] = 0
            }
        }
    }

    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j > 0; j-- {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
        if colFlag { matrix[i][0] = 0 }
    }
}

func main() {
    matrix := [][]int {
        {0, 1, 2, 0},
        {3, 4, 5, 2},
        {1, 3, 1, 5},
    }
    setZeroes(matrix)

    for _, row := range matrix {
        fmt.Println(row)
    }
}

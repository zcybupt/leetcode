package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
    var result []int
    if matrix == nil || len(matrix) == 0 { return result }

    l, r, t, b := 0, len(matrix[0]) - 1, 0, len(matrix) - 1

    for l <= r && t <= b {
        for i := l; i <= r; i++ { result = append(result, matrix[t][i]) }
        t++
        for i := t; i <= b; i++ { result = append(result, matrix[i][r]) }
        r--
        for i := r; i >= l && t <= b; i-- { result = append(result, matrix[b][i]) }
        b--
        for i := b; i >= t && l <= r; i-- { result = append(result, matrix[i][l]) }
        l++
    }

    return result
}

func main() {
    matrix := [][]int {
        {1,  2,  3,  4},
        {5,  6,  7,  8},
        {9, 10, 11, 12},
    }

    fmt.Println(spiralOrder(matrix))
}

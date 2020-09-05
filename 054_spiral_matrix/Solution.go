package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
    var result []int
    if matrix == nil || len(matrix) == 0 {
        return result
    }
    row1, col1 := 0, 0
    row2, col2 := len(matrix) - 1, len(matrix[0]) - 1

    for row1 <= row2 && col1 <= col2 {
        result = append(result, getEdge(matrix, row1, col1, row2, col2)...)
        row1, col1, row2, col2 = row1 + 1, col1 + 1, row2 - 1, col2 - 1
    }

    return result
}

func getEdge(m [][]int, row1, col1, row2, col2 int) []int {
    var tmpResult []int
    if row1 == row2 {
        for i := col1; i <= col2; i++ {
            tmpResult = append(tmpResult, m[row1][i])
        }
        return tmpResult
    }
    if col1 == col2 {
        for i := row1; i <= row2; i++ {
            tmpResult = append(tmpResult, m[i][col1])
        }
        return tmpResult
    }

    for i := col1; i < col2; i++ {
        tmpResult = append(tmpResult, m[row1][i])
    }
    for i := row1; i < row2; i++ {
        tmpResult = append(tmpResult, m[i][col2])
    }
    for i := col2; i > col1; i-- {
        tmpResult = append(tmpResult, m[row2][i])
    }
    for i := row2; i > row1; i-- {
        tmpResult = append(tmpResult, m[i][col1])
    }

    return tmpResult
}

func main() {
    nums := [][]int {
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    fmt.Println(spiralOrder(nums))
}


package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
    if matrix == nil || len(matrix) == 0 { return false }
    m, n := len(matrix), len(matrix[0])
    l, r := 0, m * n - 1

    for l <= r {
        mid := (l + r) / 2
        midVal := matrix[mid / n][mid % n]
        if target == midVal {
            return true
        } else if target > midVal {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }

    return false
}

func main() {
    matrix := [][]int {
        {1,   3,  5,  7},
        {10, 11, 16, 20},
        {23, 30, 34, 50},
    }

    fmt.Println(searchMatrix(matrix, 10))
}

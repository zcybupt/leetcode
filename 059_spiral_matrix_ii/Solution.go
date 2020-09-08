package main

import "fmt"

func generateMatrix(n int) [][]int {
    matrix := make([][]int, n)
    for i := 0; i < n; i++ {
        matrix[i] = make([]int, n)
    }

    l, r, t, b := 0, n - 1, 0, n - 1
    num := 1

    for num <= n * n {
        for i := l; i <= r; i++ {
            matrix[t][i] = num
            num++
        }
        t++
        for i := t; i <= b; i++ {
            matrix[i][r] = num
            num++
        }
        r--
        for i := r; i >= l; i-- {
            matrix[b][i] = num
            num++
        }
        b--
        for i := b; i >= t; i-- {
            matrix[i][l] = num
            num++
        }
        l++
    }

    return matrix
}

func main() {
    matrix := generateMatrix(3)
    for _, row := range matrix {
        fmt.Println(row)
    }
}

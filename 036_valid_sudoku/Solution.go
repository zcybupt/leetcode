package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
    rows := [9][9]int{}
    columns := [9][9]int{}
    boxes := [9][9]int{}

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            num := board[i][j]
            if num == '.' {
                continue
            }
            boxIndex := (i/3)*3 + j/3
            num = num - '1'
            rows[i][num]++
            columns[j][num]++
            boxes[boxIndex][num]++

            if rows[i][num] > 1 || columns[j][num] > 1 || boxes[boxIndex][num] > 1 {
                return false
            }
        }
    }

    return true
}

func main() {
    board := [][]byte{
        {'5', '3', '.', '.', '7', '.', '.', '.', '.'},
        {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
        {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
        {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
        {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
        {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
        {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
        {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
        {'.', '.', '.', '.', '8', '.', '.', '7', '9'},
    }
    fmt.Println(isValidSudoku(board))
}

package main

import "fmt"

var m, n int
var tmpBoard [][]byte
var tmpWord string
var wordLen int
var visited [][]bool
var directions = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func exist(board [][]byte, word string) bool {
    m, n = len(board), len(board[0])
    tmpBoard = board
    tmpWord = word
    wordLen = len(word)
    visited = make([][]bool, m)
    for i := 0; i < m; i++ {
        visited[i] = make([]bool, n)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if dfs(i, j, 0) {
                return true
            }
        }
    }

    return false
}

func dfs(i, j, strIdx int) bool {
    if strIdx == wordLen - 1 {
        return tmpBoard[i][j] == tmpWord[strIdx]
    }

    if tmpBoard[i][j] != tmpWord[strIdx] { return false }

    visited[i][j] = true
    for _, direction := range directions {
        newI := i + direction[0]
        newJ := j + direction[1]
        if inArea(newI, newJ) && !visited[newI][newJ] {
            if dfs(newI, newJ, strIdx + 1) {
                return true
            }
        }
    }
    visited[i][j] = false

    return false
}

func inArea(i, j int) bool {
    return 0 <= i && i < m && 0 <= j && j < n
}

func main() {
    board := [][]byte {
        {'A', 'B', 'C', 'E'},
        {'S', 'F', 'C', 'S'},
        {'A', 'D', 'E', 'E'},
    }

    fmt.Println(exist(board, "ABCCE"))
    fmt.Println(exist(board, "SEE"))
    fmt.Println(exist(board, "ABCB"))
}

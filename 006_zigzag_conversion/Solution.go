package main

import (
    "fmt"
    "strings"
)

func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }

    rows := make([]string, numRows)
    curRow := 0
    downward := false

    for _, v := range s {
        rows[curRow] += string(v)
        if curRow == 0 || curRow == numRows-1 {
            downward = !downward
        }
        if downward {
            curRow += 1
        } else {
            curRow -= 1
        }
    }

    return strings.Join(rows, "")
}

func main() {
    fmt.Println(convert("LEETCODEISHIRING", 3))
}

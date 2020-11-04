package main

import (
    "fmt"
    "strings"
)

func reverseWords(s string) string {
    s = strings.Trim(s, " ")
    l , r:= len(s) - 1, len(s) - 1
    var result []string
    for l >= 0 {
        for l >= 0 && s[l] != ' ' { l-- }
        result = append(result, s[l + 1: r + 1])
        for l >= 0 && s[l] == ' ' { l-- }
        r = l
    }

    return strings.Join(result, " ")
}

func main() {
    s := "Stay hungry, stay foolish."
    fmt.Println(reverseWords(s))
}

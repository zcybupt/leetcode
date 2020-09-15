package main

import (
    "fmt"
    "strconv"
    "strings"
)

func addBinary(a string, b string) string {
    aLen, bLen := len(a) - 1, len(b) - 1
    var result []string
    carry := 0

    for aLen >= 0 || bLen >= 0 || carry != 0 {
        curA := getNum(a, aLen)
        curB := getNum(b, bLen)
        tmp := curA + curB + carry
        carry = tmp / 2
        result = append(result, strconv.Itoa(tmp % 2))
        aLen--
        bLen--
    }

    for i, j := 0, len(result) - 1; i < j; i, j = i + 1, j - 1 {
        result[i], result[j] = result[j], result[i]
    }

    return strings.Join(result, "")
}

func getNum(s string, i int) int {
    if i >= 0 {
        return int(s[i] - '0')
    }
    return 0
}

func main() {
    fmt.Println(addBinary("11", "1"))
}

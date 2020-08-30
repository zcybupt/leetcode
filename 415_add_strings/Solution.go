package main

import (
    "fmt"
    "strconv"
    "strings"
)

func addStrings(num1 string, num2 string) string {
    i, j := len(num1)-1, len(num2)-1
    var result []string
    carry := 0

    for i >= 0 || j >= 0 || carry != 0 {
        x := getNum(num1, i)
        y := getNum(num2, j)
        tmp := x + y + carry
        result = append(result, strconv.Itoa(tmp%10))
        carry = tmp / 10
        i--
        j--
    }

    for i, j = 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }

    return strings.Join(result, "")
}

func getNum(num string, index int) int {
    if index >= 0 {
        return int(num[index] - '0')
    } else {
        return 0
    }
}

func main() {
    fmt.Println(addStrings("123", "98"))
}


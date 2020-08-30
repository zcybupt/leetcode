package main

import (
    "fmt"
    "strconv"
    "strings"
)

func multiply(num1 string, num2 string) string {
    var result string
    n1, n2 := len(num1) - 1, len(num2) - 1
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    carry := 0

    for j := n2; j >= 0; j-- {
        curNum2 := getNum(num2, j)
        var tmpRes []string
        for i := n1; i >= 0 || carry != 0; i-- {
            curNum1 := getNum(num1, i)
            tmp := curNum1 * curNum2 + carry
            carry = tmp / 10
            tmpRes = append(tmpRes, strconv.Itoa(tmp % 10))
        }
        result = addStrings(result, strings.Join(reverseArr(tmpRes), ""), n2 - j)
    }

    return result
}

func addStrings(s1, s2 string, digitCount int) string {
    for i := 0; i < digitCount; i++ {
        s2 += "0"
    }
    n1, n2 := len(s1) - 1, len(s2) - 1;
    var result []string
    carry := 0

    for n1 >= 0 || n2 >= 0 || carry != 0 {
        x := getNum(s1, n1)
        y := getNum(s2, n2)
        tmp := x + y + carry
        carry = tmp / 10
        result = append(result, strconv.Itoa(tmp % 10))
        n1--
        n2--
    }

    return strings.Join(reverseArr(result), "")
}

func reverseArr(s []string) []string {
    for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
        s[i], s[j] = s[j], s[i]
    }

    return s
}

func getNum(num string, index int) int {
    if index >= 0 {
        return int(num[index] - '0')
    } else {
        return 0
    }
}

func main() {
    fmt.Println(multiply("123", "45"))
}

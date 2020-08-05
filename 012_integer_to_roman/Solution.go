package main

import "fmt"

func intToRoman(num int) string {
    values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

    result := ""
    for i, v := range values {
        for num >= 0 && num >= v {
            num -= v
            result += symbols[i]
        }
    }

    return result
}

func main() {
    fmt.Println(intToRoman(140))
}

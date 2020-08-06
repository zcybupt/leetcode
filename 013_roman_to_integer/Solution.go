package main

import "fmt"

func romanToInt(s string) int {
    mapping := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

    result := 0
    num1 := mapping[string(s[0])]
    num2 := 0
    for _, v := range s[1:] {
        num2 = mapping[string(v)]
        if num1 >= num2 {
            result += num1
        } else {
            result -= num1
        }
        num1 = num2
    }
    result += num1

    return result
}

func main() {
    fmt.Println(romanToInt("MCMXCIV"))
}

package main

import (
    "fmt"
    "math"
)

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    result := 0.0
    maxInt := math.Pow(2, 31) - 1
    y := x

    for y != 0 {
        result = result * 10 + float64(y % 10)
        if result > maxInt {
            return false
        }
        y /= 10
    }

    return int(result) == x
}

func main() {
    fmt.Println(isPalindrome(12321))
}

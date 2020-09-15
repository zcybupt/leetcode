package main

import "fmt"

func plusOne(digits []int) []int {
    for i := len(digits) - 1; i >= 0; i-- {
        digits[i] = (digits[i] + 1) % 10
        if (digits[i] != 0) {
            return digits
        }
    }

    result := make([]int, len(digits) + 1)
    result[0] = 1
    return result
}

func main() {
    digits := []int {1, 2, 3, 4}
    result := plusOne(digits)
    for _, digit := range result {
        fmt.Printf("%d ", digit)
    }
}

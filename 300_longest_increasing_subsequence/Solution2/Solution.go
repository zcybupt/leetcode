package main

import "fmt"

func lengthOfLIS(nums []int) int {
    arrLen := len(nums)
    if arrLen < 2 { return arrLen }

    tails := make([]int, arrLen)
    result := 0
    for _, num := range nums {
        l, r := 0, result;
        for l < r {
            m := l + (r - l) / 2
            if tails[m] < num {
                l = m + 1
            } else {
                r = m
            }
        }
        tails[l] = num
        if result == r { result++ }
    }

    return result
}

func main() {
    nums := []int{10, 9, 2, 5, 3, 7, 21, 18}
    fmt.Println(lengthOfLIS(nums))
}

package main

import "fmt"

func majorityElement(nums []int) int {
    votes, x := 0, 0
    for _, num := range nums {
        if votes == 0 { x = num }
        if x == num {
            votes++
        } else {
            votes--
        }
    }

    return x
}

func main() {
    fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

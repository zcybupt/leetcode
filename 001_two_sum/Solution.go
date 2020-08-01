package main

import "fmt"

func main() {
    input := []int{3, 7, 11, 15}
    fmt.Printf("%#v", twoSum(input, 10))
}

func twoSum(nums []int, target int) []int {
    if len(nums) == 0 {
        return nil
    }

    resMap := make(map[int]int)
    for i, v := range nums {
        index, ok := resMap[v]
        if ok {
            return []int{index, i}
        } else {
            resMap[target - v] = i
        }
    }
    return nil
}

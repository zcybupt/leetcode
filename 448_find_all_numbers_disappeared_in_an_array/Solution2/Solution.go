package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
    if nums == nil || len(nums) == 0 { return []int{} }

    for _, num := range nums {
        nums[abs(num) - 1] = -abs(nums[abs(num) - 1])
    }
    var result []int
    for i, num := range nums {
        if num > 0 { result = append(result, i + 1) }
    }

    return result
}

func abs(num int) int {
    if num >= 0 { return num }
    return -num
}

func main() {
    nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
    fmt.Println(findDisappearedNumbers(nums))
}

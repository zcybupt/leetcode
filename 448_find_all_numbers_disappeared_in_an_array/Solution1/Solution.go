package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
    if nums == nil || len(nums) == 0 { return []int{} }
    tmpArr := make([]int, len(nums))

    var result []int
    for _, num := range nums { tmpArr[num - 1] += 1 }
    for i, num := range tmpArr {
        if num == 0 { result = append(result, i + 1) }
    }

    return result
}

func main() {
    nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
    fmt.Println(findDisappearedNumbers(nums))
}

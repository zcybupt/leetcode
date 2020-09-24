package main

import "fmt"

func removeDuplicates(nums []int) int {
    arrLen := len(nums)
    if arrLen <= 2 { return arrLen }

    i, count := 0, 1
    for j := 1; j < arrLen; j++ {
        if nums[i] == nums[j] {
            count++
        } else {
            count = 1
        }

        if count <= 2 {
            i++
            nums[i] = nums[j]
        }
    }

    return i + 1
}

func main() {
    nums := []int{1, 1, 1, 2, 2, 3}
    fmt.Println(removeDuplicates(nums))
    fmt.Println(nums)
}

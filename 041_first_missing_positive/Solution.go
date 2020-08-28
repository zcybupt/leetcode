package main

import "fmt"

func firstMissingPositive(nums []int) int {
    arrLen := len(nums)

    for i, _ := range nums {
        for nums[i] > 0 && nums[i] <= arrLen && nums[i] != nums[nums[i] - 1] {
            nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
        }
    }

    for i := 0; i < arrLen; i++ {
        if nums[i] != i + 1 {
            return i + 1
        }
    }

    return arrLen + 1
}

func main() {
    nums := []int{3, 4, -1, 1}
    fmt.Println(firstMissingPositive(nums))
}

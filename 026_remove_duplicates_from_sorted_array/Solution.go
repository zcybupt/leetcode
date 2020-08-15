package main

import "fmt"

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    j := 0
    for i := 0; i < len(nums); i++ {
        if (nums[i] != nums[j]) {
            j++
            nums[j] = nums[i]
        }
    }

    return j + 1
}

func main() {
    nums := []int{1, 1, 2, 2}
    newLen := removeDuplicates(nums)
    for i := 0; i < newLen; i++ {
        fmt.Printf("%d ", nums[i])
    }
}

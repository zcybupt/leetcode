package main

import "fmt"

func removeElement(nums []int, val int) int {
    j := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[j] = nums[i]
            j++
        }
    }

    return j
}

func main() {
    nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
    val := 2
    newLen := removeElement(nums, val)
    for i := 0; i < newLen; i++ {
        fmt.Printf("%d ", nums[i])
    }
}

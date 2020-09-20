package main

import "fmt"

func sortColors(nums []int)  {
    L, R, cur := 0, len(nums) - 1, 0

    for cur <= R {
        if nums[cur] == 0 {
            nums[cur], nums[L] = nums[L], nums[cur]
            cur++
            L++
        } else if nums[cur] == 2 {
            nums[cur], nums[R] = nums[R], nums[cur]
            R--
        } else {
            cur++
        }
    }
}

func main() {
    nums := []int{2, 0, 2, 1, 1, 0}
    sortColors(nums)
    fmt.Println(nums)
}

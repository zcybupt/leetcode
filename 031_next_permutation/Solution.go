package main

import "fmt"

func nextPermutation(nums []int)  {
    arrLen := len(nums)
    if arrLen <= 1 {
        return
    }

    i, j, k := arrLen - 2, arrLen - 1, arrLen - 1
    for i >= 0 && nums[i] >= nums[j] {
        i--
        j--
    }

    if i >= 0 {
        for nums[i] >= nums[k] {
            k--
        }
        nums[i], nums[k] = nums[k], nums[i]
    }

    i, j = j, arrLen - 1
    for i < j {
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6}
    nextPermutation(nums)
    fmt.Printf("%#v\n", nums)
}

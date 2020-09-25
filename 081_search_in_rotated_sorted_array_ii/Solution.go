package main

import "fmt"

func search(nums []int, target int) bool {
    arrLen := len(nums)
    l, r := 0, arrLen - 1

    for l <= r {
        for l < r && nums[l] == nums[l + 1] { l++ }
        for l < r && nums[r] == nums[r - 1] { r-- }
        mid := (l + r) / 2
        if nums[mid] == target { return true }
        if nums[l] <= nums[mid] {
            if nums[l] <= target && target < nums[mid] {
                r = mid - 1
            } else {
                l = mid + 1
            }
        } else {
            if nums[mid] < target && target <= nums[r] {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }

    return false
}

func main() {
    nums := []int {4, 5, 6, 7, 0, 1, 2}
    fmt.Println(search(nums, 2))
}

package main

import "fmt"

func search(nums []int, target int) int {
    arrLen := len(nums)
    left, right := 0, arrLen - 1

    for left <= right {
        mid := (left + right) / 2
        if nums[mid] == target {
            return mid
        }

        if nums[mid] < nums[right] {    // 右侧有序
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        } else {    // 左侧有序
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        }
    }

    return -1
}

func main() {
    nums := []int{4, 5, 6, 7, 0, 1, 2}
    target := 0
    fmt.Println(search(nums, target))
}

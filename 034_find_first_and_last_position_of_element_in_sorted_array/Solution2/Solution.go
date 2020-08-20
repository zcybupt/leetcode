package main

import "fmt"

func searchRange(nums []int, target int) []int {
    return []int {
        binarySearch(nums, target, true),
        binarySearch(nums, target, false),
    }
}

func binarySearch(nums []int, target int, isSearchFirst bool) int {
    arrLen := len(nums)
    if arrLen == 0 {
        return -1
    }
    left, right := 0, arrLen - 1

    for left <= right {
        mid := (left + right) / 2
        if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            if isSearchFirst {
                if mid > 0 && nums[mid] == nums[mid - 1] {
                    right = mid -1
                } else {
                    return mid
                }
            } else {
                if mid < arrLen - 1 && nums[mid] == nums[mid + 1] {
                    left = mid + 1
                } else {
                    return mid
                }
            }
        }
    }

    return -1
}

func main() {
    nums := []int{5, 7, 7, 8, 8, 10}
    target := 8
    fmt.Printf("%#v\n", searchRange(nums, target))
}


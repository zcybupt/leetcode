package main

import "fmt"

func searchRange(nums []int, target int) []int {
    arrLen := len(nums)
    if arrLen == 0 {
        return []int{-1, -1}
    }
    left, right := 0, arrLen - 1

    for left <= right {
        mid := (left + right) / 2
        if nums[mid] == target {
            for nums[left] != target {
                left++
            }
            for nums[right] != target {
                right--
            }
            return []int{left, right}
        }

        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return []int{-1, -1}
}

func main() {
    nums := []int{5, 7, 7, 8, 8, 10}
    target := 8
    fmt.Printf("%#v\n", searchRange(nums, target))
}

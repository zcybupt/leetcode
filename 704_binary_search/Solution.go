package main

import "fmt"

func search(nums []int, target int) int {
    if len(nums) == 0 { return -1 }

    l, r := 0, len(nums) - 1
    for l <= r {
        mid := l + (r - l) / 2
        if (nums[mid] == target) {
            return mid
        } else if (nums[mid] < target) {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }

    return -1
}

func main() {
    nums := []int{-1, 0, 3, 5, 9, 12}
    fmt.Println(search(nums, 9))
}

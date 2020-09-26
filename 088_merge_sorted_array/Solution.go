package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
    idx := m + n - 1
    m, n = m - 1, n - 1
    for m >= 0 && n >= 0 {
        if nums1[m] > nums2[n] {
            nums1[idx] = nums1[m]
            idx, m = idx - 1, m - 1
        } else {
            nums1[idx] = nums2[n]
            idx, n = idx - 1, n - 1
        }
    }

    for n >= 0 {
        nums1[idx] = nums2[n]
        idx, n = idx - 1, n - 1
    }
}

func main() {
    nums1 := []int{1, 2, 3, 0, 0, 0}
    nums2 := []int{2, 5, 6}
    m, n := 3, 3
    merge(nums1, m, nums2, n)
    fmt.Println(nums1)
}

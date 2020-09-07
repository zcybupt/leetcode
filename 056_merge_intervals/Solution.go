package main

import "fmt"

func merge(intervals [][]int) [][]int {
    if intervals == nil || len(intervals) <= 1 { return intervals }
    quickSort(intervals, 0, len(intervals) - 1)

    var result [][]int
    result = append(result, intervals[0])
    for _, interval := range intervals[1:] {
        lastInterval := result[len(result) - 1]
        if lastInterval[1] < interval[0] {
            result = append(result, interval)
        } else {
            lastInterval[1] = max(lastInterval[1], interval[1])
        }
    }

    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func quickSort(nums [][]int, L, R int) {
    if L < R {
        pIndex := partition(nums, L, R)
        quickSort(nums, L, pIndex - 1)
        quickSort(nums, pIndex + 1, R)
    }
}

func partition(nums [][]int, L, R int) int {
    P := nums[L]
    for L != R {
        for L < R && nums[R][0] >= P[0] { R-- }
        nums[L] = nums[R]
        for L < R && nums[L][0] <= P[0] { L++ }
        nums[R] = nums[L]
    }
    nums[L] = P

    return L
}

func main() {
    intervals := [][]int{{2, 6}, {1, 3}, {9, 13}, {15, 18}, {8, 10}};
    result := merge(intervals)
    fmt.Println(result)
}

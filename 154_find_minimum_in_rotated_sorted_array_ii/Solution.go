package main

import "fmt"

func minArray(nums []int) int {
    l, r := 0, len(nums) - 1
    for l <= r {
        mid := l + (r - l) / 2
        if (nums[r] > nums[mid]) {
            // 右侧有序, mid 有可能是最小值, 故下一轮应该包含 mid
            // 下一轮循环范围为 [left, mid]
            r = mid;
        } else if (nums[r] < nums[mid]) {
            // 左侧有序, mid 及 mid 左侧一定不是最小数字
            // 下一轮循环范围为 [mid + 1, right]
            l = mid + 1;
        } else {
            // mid 和 r 相等, 去重
            r--;
        }
    }

    return nums[l]
}

func main() {
    nums := []int{3, 4, 5, 1, 2}
    fmt.Println(minArray(nums))
}

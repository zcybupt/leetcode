package main

import (
    "fmt"
    "sort"
)

func threeSum(nums []int) [][]int {
    arrLen := len(nums)
	var results [][]int
	if nums == nil || arrLen < 3 {
		return results
	}

	sort.Ints(nums)

	for i := 0; i < arrLen; i++ {
		if nums[i] > 0 {
			return results
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		L, R := i+1, arrLen-1
		for L < R {
			if nums[i]+nums[L]+nums[R] == 0 {
				results = append(results, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if nums[i]+nums[L]+nums[R] > 0 {
				R--
			} else {
				L++
			}
		}
	}

	return results
}

func main() {
    nums := []int{-1, 0, 1, 2, -1, -4}
    fmt.Printf("%#v\n", threeSum(nums))
}

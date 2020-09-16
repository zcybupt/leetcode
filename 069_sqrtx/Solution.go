package main

import "fmt"

func mySqrt(x int) int {
    l, r, res := 0, x, -1

    for l <= r {
        mid := l + (r - l) / 2
        if mid * mid <= x {
            res = mid
            l = mid + 1
        } else {
            r = mid - 1
        }
    }

    return res
}

func main() {
    fmt.Println(mySqrt(5))
}

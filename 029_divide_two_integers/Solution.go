package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	sign := 1
	if (dividend > 0) != (divisor > 0) {
		sign = -1
	}
	m, n := abs(dividend), abs(divisor)
	res := 0

	for m >= n {
		t, p := n, 1
		for m > (t << 1) {
			t <<= 1
			p <<= 1
		}
		m -= t
		res += p
	}

	if sign == 1 {
		return int(math.Min(float64(res), math.MaxInt32))
	} else {
		return int(math.Max(float64(-res), math.MinInt32))
	}
}

func abs(x int) int64 {
	if x > 0 {
		return int64(x)
	} else {
		return -int64(x)
	}
}

func main() {
	fmt.Println(divide(-2147483648, 1))
}


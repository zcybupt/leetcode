package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	result := 0.0
	rightLimit := math.Pow(2, 31) - 1
	leftLimit := -rightLimit - 1

	for x != 0 {
		result = result*10.0 + float64(x%10.0)
		if result > rightLimit || result < leftLimit {
			return 0
		}
		x /= 10.0
	}

	return int(result)
}

func main() {
	fmt.Println(reverse(-123))
}


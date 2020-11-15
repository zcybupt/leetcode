package main

import (
    "fmt"
    "math"
)

func maxProfit(prices []int) int {
    cost, profit := math.MaxInt32, 0
    for _, price := range prices {
        cost = min(cost, price)
        profit = max(profit, price - cost)
    }

    return profit
}

func max(a, b int) int {
    if a > b { return a }
    return b
}

func min(a, b int) int {
    if a < b { return a }
    return b
}

func main() {
    prices := []int{7, 1, 5, 3, 6, 4}
    fmt.Println(maxProfit(prices))
}

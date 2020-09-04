package main

import "fmt"

func quickMul(x float64, n int) float64 {
    result := 1.0
    for n > 0 {
        if n % 2 == 1 {
            result *= x
        }
        x *= x
        n /= 2
    }
    return result
}

func myPow(x float64, n int) float64 {
    if n >= 0 {
        return quickMul(x, n)
    }
    return 1 / quickMul(x, -n)
}

func main() {
    fmt.Println(myPow(2, 10))
}

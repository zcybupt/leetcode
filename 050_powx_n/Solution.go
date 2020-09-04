package main

import "fmt"

func quickMul(x float64, n int) float64 {
    if n == 0 {
        return 1.0
    }

    y := quickMul(x, n / 2)

    if n % 2 == 0 {
        return y * y
    }
    return y * y * x
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

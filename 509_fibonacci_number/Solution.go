package main

import "fmt"

func fib(n int) int {
    if n < 2 { return n }

    a, b := 0, 1
    for n > 1 {
        a, b = b, a + b
        n--
    }

    return b
}

func main() {
    fmt.Println(fib(7))
}

package main

import "fmt"

func grayCode(n int) []int {
    var result []int
    for i := 0; i < 1<<n; i++ {
        result = append(result, i ^ i>>1)
    }
    return result
}

func main() {
    fmt.Println(grayCode(3))
}

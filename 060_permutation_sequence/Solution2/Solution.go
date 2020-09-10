package main

import (
    "fmt"
    "strconv"
    "strings"
)

var facts = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

func getFact(n int) int {
    if n < 10 {
        return facts[n]
    }
    return n * getFact(n-1)
}

func getPermutation(n int, k int) string {
    var result []string
    isVisited := make([]bool, n)
    k--
    for i := 0; i < n; i++ {
        tmpFact := getFact(n - 1 - i)
        cnt := k / tmpFact
        k %= tmpFact
        for j := 0; j < n; j++ {
            if isVisited[j] {
                continue
            }
            if cnt == 0 {
                isVisited[j] = true
                result = append(result, strconv.Itoa(j+1))
                break
            }
            cnt--
        }
    }

    return strings.Join(result, "")
}

func main() {
    fmt.Println(getPermutation(4, 9))
}


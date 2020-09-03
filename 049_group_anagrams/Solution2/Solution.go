package main

import (
    "fmt"
)

func groupAnagrams(strs []string) [][]string {
    primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
    tmpMap := make(map[int][]string)
    for _, str := range strs {
        key := 1
        for _, ch := range str {
            key *= primes[ch-'a']
        }
        if _, ok := tmpMap[key]; !ok {
            tmpMap[key] = []string{str}
        } else {
            tmpMap[key] = append(tmpMap[key], str)
        }
    }

    results := make([][]string, len(tmpMap))
    i := 0
    for _, result := range tmpMap {
        results[i] = result
        i++
    }

    return results
}

func main() {
    strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
    results := groupAnagrams(strs)
    for _, v := range results {
        fmt.Println(v)
    }
}


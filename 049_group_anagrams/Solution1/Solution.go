package main

import (
    "fmt"
    "sort"
    "strings"
)

func groupAnagrams(strs []string) [][]string {
    tmpMap := make(map[string][]string)
    for _, str := range strs {
        chars := strings.Split(str, "")
        sort.Strings(chars)
        key := strings.Join(chars, "")
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


package main

import "fmt"

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    commonPrefix := strs[0]
    for i, v := range strs[1:] {
        j := 0
        for ;j < len(commonPrefix) && j < len(v); j++ {
            if strs[i][j] != commonPrefix[j] {
                break
            }
        }
        commonPrefix = commonPrefix[:j]
        if len(commonPrefix) == 0 {
            return ""
        }
    }

    return commonPrefix
}

func main() {
    strs := []string{"flower", "flow", "flour"}
    fmt.Println(longestCommonPrefix(strs))
}

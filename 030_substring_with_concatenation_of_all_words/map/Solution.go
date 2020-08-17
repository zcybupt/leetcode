package main

import "fmt"

func findSubstring(s string, words []string) []int {
    var results []int
    sLen := len(s)
    arrLen := len(words)
    if sLen == 0 || arrLen == 0 {
        return results
    }
    wordLen := len(words[0])

    wordMap := make(map[string]int)
    for _, key := range words {
        wordMap[key] = wordMap[key] + 1
    }

    for i := 0; i < sLen - wordLen * arrLen + 1; i++ {
        curMap := make(map[string]int)
        count := 0
        for count < arrLen {
            curWord := s[i + wordLen * count: i + wordLen * (count + 1)]
            if wordMap[curWord] != 0 {
                curMap[curWord]++
                if curMap[curWord] > wordMap[curWord] {
                    break
                }
            } else {
                break
            }
            count++
        }
        if count == arrLen {
            results = append(results, i)
        }
    }
    return results
}

func main() {
    s := "barfoothefoobarman"
    words := []string{"bar", "foo"}
    results := findSubstring(s, words)
    for _, result := range results {
        fmt.Printf("%d ", result)
    }
}

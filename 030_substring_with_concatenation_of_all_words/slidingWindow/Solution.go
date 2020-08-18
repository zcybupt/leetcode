package main

import "fmt"

func findSubstring(s string, words []string) []int {
    var results []int
    sLen, arrLen := len(s), len(words)
    if sLen == 0 || arrLen == 0 {
        return results
    }
    wordLen := len(words[0])

    wordMap := make(map[string]int)
    for _, word := range words {
        wordMap[word]++
    }

    windowSize := wordLen * arrLen;
    for i := 0; i < wordLen; i++ {
        start := i
        for start + windowSize <= sLen {
            curStr := s[start: start + windowSize]
            curMap := make(map[string]int)
            j := arrLen
            for j > 0 {
                curWord := curStr[(j-1) * wordLen: j * wordLen]
                count := curMap[curWord] + 1
                if count > wordMap[curWord] {
                    break
                }
                curMap[curWord] = count
                j--
            }
            if j == 0 {
                results = append(results, start)
            }
            start += wordLen * max(j, 1)
        }
    }

    return results
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func main() {
    s := "barfoothefoobarman"
    words := []string{"bar", "foo"}
    results := findSubstring(s, words)
    for _, result := range results {
        fmt.Printf("%d ", result)
    }
}

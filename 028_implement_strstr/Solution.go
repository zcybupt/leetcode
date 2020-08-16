package main

import "fmt"

func getNextTable(needle string) []int {
    needleLen := len(needle)
    nextTable := make([]int, needleLen)
    nextTable[0] = -1
    lp, rp := -1, 0

    for rp < needleLen - 1 {
        if lp == -1 ||needle[lp] == needle[rp] {
            lp++
            rp++
            nextTable[rp] = lp
        } else {
            lp = nextTable[lp]
        }
    }

    return nextTable
}

func strStr(haystack string, needle string) int {
    haystackLen := len(haystack)
    needleLen := len(needle)
    if needleLen == 0 {
        return 0
    }

    nextTable := getNextTable(needle)
    haystackCur, needleCur := 0, 0
    for haystackCur < haystackLen && needleCur < needleLen {
        if haystack[haystackCur] == needle[needleCur] {
            haystackCur++
            needleCur++
        } else {
            if needleCur == 0 {
                haystackCur++
            } else {
                needleCur = nextTable[needleCur]
            }
        }
    }

    if needleCur == needleLen {
        return haystackCur - needleCur
    } else {
        return -1
    }
}

func main() {
    fmt.Println(strStr("hello", "ll"))
}

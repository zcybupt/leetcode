package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

func simplifyPath(path string) string {
    return filepath.Clean(path)
}

func simplifyPath2(path string) string {
    pathArr := strings.Split(path, "/")
    var stack []string
    for _, tmp := range pathArr {
        if tmp == "" || tmp == "." { continue }
        if tmp == ".." {
            if len(stack) != 0 {
                stack = stack[:len(stack) - 1]
            }
            continue
        }
        stack = append(stack, tmp)
    }

    return "/" + strings.Join(stack, "/")
}

func max(a, b int) int {
    if a > b { return a }
    return b
}

func main() {
    testCases := []string {
        "/...",
        "/home/",
        "/../",
        "/home//foo/",
        "/a/../../b/../c//.//",
        "/a//b////c/d//././/..",
        "/a/./b/../../c/",
        "/a//b////c/d//././/..",
    }

    for _, test := range testCases {
        fmt.Println(simplifyPath2(test))
    }
}

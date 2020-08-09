package main

import "fmt"

var mapping = map[string][]string {
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

var result []string

func letterCombinations(digits string) []string {
	result = []string{}
	if len(digits) == 0 {
		return result
	}

	dfs(digits, "")
	return result
}

func dfs(digits, tmpStr string) {
	if len(digits) == 0 {
		result = append(result, tmpStr)
		return
	}

	for _, v := range mapping[digits[0:1]] {
		dfs(digits[1:], tmpStr + v)
	}
}

func main() {
	fmt.Printf("%#v\n", letterCombinations("234"))
}

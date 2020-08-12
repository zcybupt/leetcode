package main

import "fmt"

func generateParenthesis(n int) []string {
	var result []string
	var dfs func(l, r int, curStr string)

	dfs = func(l, r int, curStr string) {
		if l == 0 && r == 0 {
			result = append(result, curStr)
			return
		}

		if l > 0 {
			dfs(l-1, r, curStr+"(")
		}
		if r > l {
			dfs(l, r-1, curStr+")")
		}
	}
	dfs(n, n, "")

	return result
}

func main() {
	result := generateParenthesis(3)
	for _, v := range result {
		fmt.Println(v)
	}
}


package main

import "fmt"

func isValid(s string) bool {
	mapping := map[int32]int32{')': '(', ']': '[', '}': '{'}
	stack := make([]int32, len(s))
	index := 0

	for _, v := range s {
		if index > 0 && stack[index-1] == mapping[v] {
			index--
			continue
		}
		stack[index] = v
		index++
	}

	return index == 0
}

func main() {
	fmt.Println(isValid("()[{}]"))
}


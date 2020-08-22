package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	tmp1 := "1"
	for i := 2; i <= n; i++ {
		tmp2 := ""
		count := 1
		for j := 1; j < len(tmp1); j++ {
			if tmp1[j] == tmp1[j - 1] {
				count++
			} else {
				tmp2 += strconv.Itoa(count) + string(tmp1[j - 1])
				count = 1
			}
		}
		tmp2 += strconv.Itoa(count) + string(tmp1[len(tmp1)-1])
		tmp1 = tmp2
	}

	return tmp1
}

func main() {
	fmt.Println(countAndSay(4))
}


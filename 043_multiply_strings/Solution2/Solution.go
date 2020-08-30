package main

import (
	"fmt"
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1, n2 := len(num1), len(num2)
	tmpRes := make([]int, n1 + n2)

	for i := n1 - 1; i >= 0; i-- {
		x := num1[i] - '0'
		for j := n2 - 1; j >= 0; j-- {
			y := num2[j] - '0'
			tmp := int(x * y) + tmpRes[i + j + 1]
			tmpRes[i + j + 1] = tmp % 10
			tmpRes[i + j] += tmp / 10
		}
	}

	var result []string
	for i, v := range tmpRes {
		if i == 0 && v == 0 {
			continue
		}
		result = append(result, strconv.Itoa(v))
	}

	return strings.Join(result, "")
}

func main() {
	fmt.Println(multiply("123", "45"))
}

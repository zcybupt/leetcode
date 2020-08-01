package main

import "fmt"

func main() {
	l1 := &ListNode{2, nil}
	l1.Next = &ListNode{4, nil}
	l1.Next.Next = &ListNode{3, nil}

	l2 := &ListNode{5, nil}
	l2.Next = &ListNode{6, nil}
	l2.Next.Next = &ListNode{4, nil}

	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Printf("%v ", result.Val)
		result = result.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p, q := l1, l2
	dummyHead := &ListNode{0, nil}
	curr := dummyHead
	carry := 0

	for p != nil || q != nil {
		x, y := 0, 0
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}

		if q != nil {
			y = q.Val
		} else {
			y = 0
		}

		sum := carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next

		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}

	return dummyHead.Next
}


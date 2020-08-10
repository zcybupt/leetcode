package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if removeNode(head, n) == n {
        return head.Next
    } else {
        return head
    }
}

func removeNode(head *ListNode, n int) int {
    if head.Next == nil {
        return 1
    }

    m := removeNode(head.Next, n)
    if m == n {
        head.Next = head.Next.Next
    }

    return m + 1
}

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil}

	result := removeNthFromEnd(head, 2)
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
}


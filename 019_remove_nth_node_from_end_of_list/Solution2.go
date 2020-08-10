package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    fast, slow := head, head

    for i := 0; i < n; i++ {
        fast = fast.Next
    }

    if fast == nil {
        return head.Next
    }

    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next

    return head
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


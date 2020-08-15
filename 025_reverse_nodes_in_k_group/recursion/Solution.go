package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    tail := head
    for i := 0; i < k; i++ {
        if tail == nil {
            return head
        }
        tail = tail.Next
    }

    new_head := reverse(head, tail)
    head.Next = reverseKGroup(tail, k)

    return new_head
}

func reverse(head, tail *ListNode) *ListNode {
    var pre *ListNode

    for head != tail {
        next := head.Next
        head.Next, pre, head = pre, head, next
    }

    return pre
}

func main() {
    head := &ListNode{0, nil}
    cur := head
    for i := 0; i < 10; i++ {
        cur.Next = &ListNode{i, nil}
        cur = cur.Next
    }

    cur = reverseKGroup(head, 2)
    for cur != nil {
        fmt.Printf("%d ", cur.Val)
        cur = cur.Next
    }
}

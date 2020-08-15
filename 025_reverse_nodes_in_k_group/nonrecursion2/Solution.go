package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    dummyHead := &ListNode{-1, head}
    pre, end := dummyHead, dummyHead

    for end.Next != nil {
        for i := 0; i < k; i++ {
            end = end.Next
            if end == nil {
                return dummyHead.Next
            }
        }
        start := pre.Next
        next := end.Next
        pre.Next = reverse(start, next)
        start.Next = next
        pre = start
        end = pre
    }

    return dummyHead.Next
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
    head := &ListNode{1, nil}
    cur := head
    for i := 2; i < 10; i++ {
        cur.Next = &ListNode{i, nil}
        cur = cur.Next
    }

    cur = reverseKGroup(head, 2)
    for cur != nil {
        fmt.Printf("%d ", cur.Val)
        cur = cur.Next
    }
}

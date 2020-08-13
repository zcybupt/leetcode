package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    cur := head
    var pre *ListNode

    for cur != nil {
        cur.Next, pre, cur = pre, cur, cur.Next
    }

    return pre
}

func reverseList2(head *ListNode) *ListNode {
    return reverse(nil, head)
}

func reverse(pre, cur *ListNode) *ListNode {
    if cur == nil {
        return pre
    }

    next := cur.Next
    cur.Next = pre

    return reverse(cur, next)
}

func main() {
    head := &ListNode{0, nil}
    cur := head

    for i := 1; i < 10; i++ {
        cur.Next = &ListNode{i ,nil}
        cur = cur.Next
    }

    cur = reverseList2(head)
    for cur != nil {
        fmt.Printf("%d ", cur.Val)
        cur = cur.Next
    }
}

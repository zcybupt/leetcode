package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil { return head }

    dummyHead := &ListNode{-1, nil}
    tail := dummyHead
    for l, r := head, head; l != nil; l = r {
        for r != nil && l.Val == r.Val { r = r.Next }
        if l.Next == r {
            tail.Next = l
            tail = l
            tail.Next = nil
        }
    }

    return dummyHead.Next
}

func main() {
    head := &ListNode{1, nil}
    head.Next = &ListNode{1, nil}
    head.Next.Next = &ListNode{2, nil}
    head.Next.Next.Next = &ListNode{3, nil}
    head.Next.Next.Next.Next = &ListNode{3, nil}

    result := deleteDuplicates(head)
    for result != nil {
        fmt.Println(result.Val)
        result = result.Next
    }
}

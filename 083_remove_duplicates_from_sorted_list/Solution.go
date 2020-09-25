package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil { return head }

    tmp := head
    for tmp != nil && tmp.Next != nil {
        if tmp.Val == tmp.Next.Val {
            tmp.Next = tmp.Next.Next
        } else {
            tmp = tmp.Next
        }
    }

    return head
}

func main() {
    head := &ListNode{1, nil}
    head.Next = &ListNode{1, nil}
    head.Next.Next = &ListNode{2, nil}
    head.Next.Next.Next = &ListNode{3, nil}
    head.Next.Next.Next.Next = &ListNode{3, nil}

    deleteDuplicates(head)

    for head != nil {
        fmt.Println(head.Val)
        head = head.Next
    }
}

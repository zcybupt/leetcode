package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil { return head }

    if head.Val == head.Next.Val {
        for head != nil && head.Next != nil && head.Val == head.Next.Val {
            head = head.Next
        }
        return deleteDuplicates(head.Next)
    } else {
        head.Next = deleteDuplicates(head.Next)
        return head
    }
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

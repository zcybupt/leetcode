package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil { return nil }
    tail := head
    length := 1
    for tail.Next != nil {
        length++
        tail = tail.Next
    }
    tail.Next = head
    for i := length - k % length; i > 0; i-- {
        tail = tail.Next
    }
    head = tail.Next
    tail.Next = nil

    return head
}

func printList(head *ListNode) {
    for head != nil {
        fmt.Printf("%d ", head.Val)
        head = head.Next
    }
}

func main() {
    head := &ListNode{1, nil}
    tmp := head
    for i := 2; i < 6; i++ {
        tmp.Next = &ListNode{i, nil}
        tmp = tmp.Next
    }
    head = rotateRight(head, 12)
    printList(head)
}


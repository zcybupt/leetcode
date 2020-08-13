package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    var latter *ListNode
    cur := head
    head = head.Next

    for cur != nil && cur.Next != nil {
        tmp := cur.Next
        cur.Next = tmp.Next
        tmp.Next = cur
        if latter != nil {
            latter.Next = tmp
        }
        latter = cur
        cur = cur.Next
    }

    return head
}

func swapPairs2(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    first, second := head, head.Next
    first.Next = swapPairs2(second.Next)
    second.Next = first

    return second
}

func main() {
    head := &ListNode{1, nil}
    cur := head

    for i := 2; i < 10; i++ {
        cur.Next = &ListNode{i, nil}
        cur = cur.Next
    }

    cur = swapPairs(head)

    for cur != nil {
        fmt.Printf("%d ", cur.Val)
        cur = cur.Next
    }
}

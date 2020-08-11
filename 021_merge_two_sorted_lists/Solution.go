package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }

    result := &ListNode{0, nil}
    cur := result

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            cur.Next = l1
            l1 = l1.Next
        } else {
            cur.Next = l2
            l2 = l2.Next
        }
        cur = cur.Next
    }

    if l1 != nil {
        cur.Next = l1
    } else {
        cur.Next = l2
    }

    return result.Next
}

func main() {
    l1 := &ListNode{1, nil}
    l1.Next = &ListNode{2, nil}
    l1.Next.Next = &ListNode{4, nil}

    l2 := &ListNode{1, nil}
    l2.Next = &ListNode{3, nil}
    l2.Next.Next = &ListNode{4, nil}

    result := mergeTwoLists(l1, l2)
    for result != nil {
        fmt.Printf("%d ", result.Val)
        result = result.Next;
    }
}

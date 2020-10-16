package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func createList(nums []int) *ListNode {
    head := &ListNode{nums[0], nil}
    tmp := head

    for _, num := range nums[1:] {
        tmp.Next = &ListNode{num, nil}
        tmp = tmp.Next
    }

    return head
}

func printList(head *ListNode) {
    for head != nil {
        fmt.Printf("%d ", head.Val)
        head = head.Next
    }
    fmt.Println()
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil { return l2 }
    if l2 == nil { return l1 }

    dummyHead := &ListNode{-1, nil}
    tmp := dummyHead
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            tmp.Next = l1
            l1 = l1.Next
        } else {
            tmp.Next = l2
            l2 = l2.Next
        }
        tmp = tmp.Next
    }

    if l1 != nil { tmp.Next = l1 }
    if l2 != nil { tmp.Next = l2 }

    return dummyHead.Next
}

func main() {
    head1 := createList([]int{1, 2, 4})
    head2 := createList([]int{1, 3, 4})
    head := mergeTwoLists(head1, head2)
    printList(head)
}

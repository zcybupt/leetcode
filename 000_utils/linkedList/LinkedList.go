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


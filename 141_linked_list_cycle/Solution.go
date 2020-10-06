package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    if head == nil { return false }

    fast, slow := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if fast == slow { return true }
    }

    return false
}

func createList(nums []int) *ListNode {
    head := &ListNode{nums[0], nil}
    tmp := head

    for _, num := range nums[1:] {
        tmp.Next = &ListNode{num, nil}
        tmp = tmp.Next
    }
    tmp.Next = head

    return head
}

func main() {
    head := createList([]int{1, 2, 3, 4, 5})
    fmt.Println(hasCycle(head))
}

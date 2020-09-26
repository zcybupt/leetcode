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

func partition(head *ListNode, x int) *ListNode {
    left, right := &ListNode{-1, nil}, &ListNode{-1, nil}

    tmp1, tmp2 := left, right
    for head != nil {
        if head.Val < x {
            tmp1.Next = head
            head = head.Next
            tmp1 = tmp1.Next
            tmp1.Next = nil
        } else {
            tmp2.Next = head
            head = head.Next
            tmp2 = tmp2.Next
            tmp2.Next = nil
        }
    }
    tmp1.Next = right.Next

    return left.Next
}

func main() {
    nums := []int{1, 4, 3, 2, 5, 2}
    head := createList(nums)
    printList(head)

    result := partition(head, 3)
    printList(result)
}

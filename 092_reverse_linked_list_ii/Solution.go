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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head == nil || head.Next == nil || m >= n { return head }

    dummyHead := &ListNode{-1, head}
    pre := dummyHead

    for i := 1; i < m; i++ { pre = pre.Next }
    if pre.Next == nil { return head }

    nodeBeforeM, nodeM, cur, next := pre, pre.Next, pre.Next, pre
    for i := m; i <= n; i++ {
        next = cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    nodeBeforeM.Next = pre
    nodeM.Next = next

    return dummyHead.Next
}

func main() {
    head := createList([]int{1, 2, 3, 4, 5})
    result := reverseBetween(head, 2, 4)
    printList(result)
}

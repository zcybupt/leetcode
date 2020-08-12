package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i].Val, (*pq)[j].Val = (*pq)[j].Val, (*pq)[i].Val
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	result := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return result
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	dummyHead := &ListNode{0, nil}
	curNode := dummyHead
	var pq PriorityQueue

	for _, list := range lists {
		tmp := list
		for tmp != nil {
			heap.Push(&pq, tmp)
			tmp = tmp.Next
		}
	}

	for len(pq) != 0 {
		curNode.Next = heap.Pop(&pq).(*ListNode)
		curNode = curNode.Next
	}
	curNode.Next = nil

	return dummyHead.Next
}

func main() {
	l1 := &ListNode{1, nil}
	l1.Next = &ListNode{4, nil}
	l1.Next.Next = &ListNode{5, nil}

	l2 := &ListNode{1, nil}
	l2.Next = &ListNode{3, nil}
	l2.Next.Next = &ListNode{4, nil}

	l3 := &ListNode{2, nil}
	l3.Next = &ListNode{6, nil}

	lists := []*ListNode{l1, l2, l3}

	result := mergeKLists(lists)
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
}


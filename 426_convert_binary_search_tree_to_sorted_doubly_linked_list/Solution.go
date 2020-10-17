package main

import "fmt"

type Node struct {
    Val   int
    Left  *Node
    Right *Node
}

func createTreeUsingBFS(nums []int) *Node {
    var nodeArr []*Node

    for i, num := range nums {
        nodeArr = append(nodeArr, &Node{num, nil, nil})
        if i == 0 { continue }
        if i % 2 == 1 {
            nodeArr[(i - 1) / 2].Left = nodeArr[i]
        } else {
            nodeArr[(i - 1) / 2].Right = nodeArr[i]
        }
    }

    return nodeArr[0]
}

func printTree(root *Node) {
    if root == nil { return }
    printTree(root.Left)
    fmt.Printf("%d ", root.Val)
    printTree(root.Right)
}

func printList(head *Node) {
    tmp := head
    for {
        fmt.Printf("%d ", tmp.Val)
        tmp = tmp.Right
        if tmp == head { break }
    }
    fmt.Println()
}

var pre, head *Node
func treeToDoublyList(root *Node) *Node {
    if root == nil { return nil }

    dfs(root)
    head.Left = pre
    pre.Right = head

    return head
}

func dfs(cur *Node) {
    if cur == nil { return }
    dfs(cur.Left)
    if pre != nil {
        pre.Right = cur
    } else {
        head = cur
    }
    cur.Left = pre
    pre = cur
    dfs(cur.Right)
}

func main() {
    root := createTreeUsingBFS([]int{10, 5, 15, 2, 6, 11, 16, 1, 3})
    printTree(root)
    fmt.Println()
    head := treeToDoublyList(root)
    printList(head)
}

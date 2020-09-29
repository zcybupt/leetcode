package main

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func createTreeUsingBFS(nums []int) *TreeNode {
    arrLen := len(nums)
    var nodeArr []*TreeNode

    for i := 0; i < arrLen; i++ {
        nodeArr = append(nodeArr, &TreeNode{nums[i], nil, nil})
        if i == 0 { continue }
        if i % 2 == 1 {
            nodeArr[(i - 1) / 2].Left = nodeArr[i]
        } else {
            nodeArr[(i - 1) / 2].Right = nodeArr[i]
        }
    }

    return nodeArr[0]
}

func dfs(root *TreeNode) {
    if root == nil { return }
    fmt.Printf("%d ", root.Val)
    dfs(root.Left)
    dfs(root.Right)
}

func bfs(root *TreeNode) {
    queue := []*TreeNode{root}

    for len(queue) != 0 {
        curNode := queue[0]
        queue = queue[1:]
        fmt.Printf("%d ", curNode.Val)
        if curNode.Left != nil { queue = append(queue, curNode.Left) }
        if curNode.Right != nil { queue = append(queue, curNode.Right) }
    }
}


func main() {
    nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
    root := createTreeUsingBFS(nums)
    bfs(root)
}

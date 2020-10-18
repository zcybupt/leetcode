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

func rightSideView(root *TreeNode) []int {
    if root == nil { return []int{} }

    var result []int
    queue := []*TreeNode{root}
    for len(queue) != 0 {
        for i := len(queue); i > 0; i-- {
            curNode := queue[0]
            queue = queue[1:]
            if i == 1 { result = append(result, curNode.Val) }
            if curNode.Left != nil { queue = append(queue, curNode.Left) }
            if curNode.Right != nil { queue = append(queue, curNode.Right) }
        }
    }

    return result
}

func main() {
    nums := []int{10, 5, 15, 2, 6, 11, 16, 1, 3}
    root := createTreeUsingBFS(nums)
    fmt.Println(rightSideView(root))
}


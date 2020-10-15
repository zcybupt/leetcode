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

func zigzagLevelOrder(root *TreeNode) [][]int {
    var result [][]int
    if root == nil { return result }

    queue := []*TreeNode{root}
    flag := true
    for len(queue) != 0 {
        var tmpRes []int
        for i := len(queue); i > 0; i-- {
            curNode := queue[0]
            queue = queue[1:]
            if flag {
                tmpRes = append(tmpRes, curNode.Val)
            } else {
                tmpRes = append(append([]int{}, curNode.Val), tmpRes...)
            }
            if curNode.Left != nil { queue = append(queue, curNode.Left) }
            if curNode.Right != nil { queue = append(queue, curNode.Right) }
        }
        result = append(result, tmpRes)
        flag = !flag
    }

    return result
}

func main() {
    root := createTreeUsingBFS([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
    result := zigzagLevelOrder(root)
    for _, level := range result {
        fmt.Println(level)
    }
}


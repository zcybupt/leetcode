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

func isSymmetric(root *TreeNode) bool {
    if root == nil { return true }
    return isMirror(root.Left, root.Right)
}

func isMirror(l, r *TreeNode) bool {
    if l == nil && r == nil { return true }
    if l == nil || r == nil || l.Val != r.Val { return false }

    return isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
}

func main() {
    root := createTreeUsingBFS([]int{1, 2, 2, 3, 4, 4, 3})
    fmt.Println(isSymmetric(root))
}

package main

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
    var result []int
    if root == nil { return result }
    var stack []*TreeNode

    for len(stack) != 0 || root != nil {
        if root != nil {
            stack = append(stack, root)
            root = root.Left
        } else {
            lastIdx := len(stack) - 1
            root = stack[lastIdx]
            result = append(result, root.Val)
            stack = stack[:lastIdx]
            root = root.Right
        }
    }

    return result
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

func main() {
    nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
    root := createTreeUsingBFS(nums)
    result := inorderTraversal(root)
    for _, num := range result {
        fmt.Printf("%d ", num)
    }
}

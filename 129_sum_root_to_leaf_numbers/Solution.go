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

func sumNumbers(root *TreeNode) int {
    if root == nil { return 0 }

    return dfs(root, 0)
}

func dfs(root *TreeNode, sum int) int {
    if root == nil { return 0 }
    if root.Left == nil && root.Right == nil { return sum + root.Val }
    sum += root.Val
    return dfs(root.Left, 10 * sum) + dfs(root.Right, 10 * sum)
}

func main() {
    nums := []int{4, 9, 0, 5, 1}
    root := createTreeUsingBFS(nums)
    fmt.Println(sumNumbers(root))
}

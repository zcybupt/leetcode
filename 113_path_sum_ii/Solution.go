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

var result [][]int
func pathSum(root *TreeNode, sum int) [][]int {
    result = result[:0]
    dfs(root, &[]int{}, sum)
    return result
}

func dfs(root *TreeNode, path *[]int, sum int) {
    if root == nil { return }
    *path = append(*path, root.Val)
    sum -= root.Val
    if sum == 0 && root.Left == nil && root.Right == nil {
        result = append(result, append([]int{}, *path...))
    }
    dfs(root.Left, path, sum)
    dfs(root.Right, path, sum)
    *path = (*path)[:len(*path) - 1]
}

func main() {
    nums := []int{5, 4, 8, 11, 13, 4, 7, 2, 5, 1}
    root := createTreeUsingBFS(nums)
    result := pathSum(root, 22)
    for _, row := range result {
        fmt.Println(row)
    }
}

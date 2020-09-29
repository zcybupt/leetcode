class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def create_tree_using_bfs(nums: list) -> TreeNode:
    arr_len = len(nums)
    node_list = []

    for i in range(arr_len):
        new_node = TreeNode(nums[i])
        node_list.append(new_node)
        if i == 0: continue
        if i % 2 == 1: node_list[(i - 1) // 2].left = new_node
        else: node_list[(i - 1) // 2].right = new_node

    return node_list[0]

class Solution:
    def sumNumbers(self, root: TreeNode) -> int:
        def dfs(root: TreeNode, sum: int) -> int:
            if not root: return 0
            if not root.left and not root.right: return sum + root.val
            sum += root.val
            return dfs(root.left, 10 * sum) + dfs(root.right, 10 * sum)

        return dfs(root, 0)


if __name__ == '__main__':
    nums = [4, 9, 0, 5, 1]
    root = create_tree_using_bfs(nums)
    print(Solution().sumNumbers(root))


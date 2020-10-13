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
    def pathSum(self, root: TreeNode, sum: int) -> list:
        result = []
        def recur(root: TreeNode, path: list, sum: int) -> None:
            if not root: return
            path.append(root.val)
            sum -= root.val
            if sum == 0 and not root.left and not root.right:
                result.append(path[:])
            recur(root.left, path, sum)
            recur(root.right, path, sum)
            path.pop()

        recur(root, [], sum)
        return result


if __name__ == '__main__':
    nums = [5, 4, 8, 11, 13, 4, 7, 2, 3, 4, 5, 1]
    root = create_tree_using_bfs(nums)
    result = Solution().pathSum(root, 22)
    for tmp in result: print(tmp)


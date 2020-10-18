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
    def rightSideView(self, root: TreeNode) -> list:
        if not root: return []

        queue = [root]
        result = []
        while queue:
            for i in range(len(queue), 0, -1):
                cur = queue[0]
                queue = queue[1:]
                if i == 1: result.append(cur.val)
                if cur.left: queue.append(cur.left)
                if cur.right: queue.append(cur.right)

        return result


if __name__ == '__main__':
    root = create_tree_using_bfs([10, 5, 15, 2, 6, 11, 16, 1, 3])
    print(Solution().rightSideView(root))


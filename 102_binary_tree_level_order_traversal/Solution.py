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
    def levelOrder(self, root: TreeNode) -> list:
        if not root: return []

        result, queue = [], [root]
        while queue:
            tmp_res = []
            for _ in range(len(queue)):
                cur = queue[0]
                queue = queue[1:]
                tmp_res.append(cur.val)
                if cur.left: queue.append(cur.left)
                if cur.right: queue.append(cur.right)
            result.append(tmp_res)
        return result

if __name__ == '__main__':
    nums = [i for i in range(14)]
    root = create_tree_using_bfs(nums)
    for tmp in Solution().levelOrder(root):
        print(tmp)


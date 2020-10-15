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
    def zigzagLevelOrder(self, root: TreeNode) -> list:
        if not root: return []

        result, queue = [], [root]
        flag = True
        while queue:
            tmp_res = []
            for _ in range(len(queue)):
                cur_node = queue[0]
                queue = queue[1:]
                if flag: tmp_res.append(cur_node.val)
                else: tmp_res = [cur_node.val] + tmp_res
                if cur_node.left: queue.append(cur_node.left)
                if cur_node.right: queue.append(cur_node.right)
            flag = not flag
            result.append(tmp_res)

        return result


if __name__ == '__main__':
    root = create_tree_using_bfs([1, 2, 3, 4, 5, 6, 7, 8, 9])
    for level in Solution().zigzagLevelOrder(root):
        print(level)


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
    def isSymmetric(self, root: TreeNode) -> bool:
        if not root: return True
        def is_mirror(l: TreeNode, r: TreeNode) -> bool:
            if not l and not r: return True
            if not l or not r or l.val != r.val: return False

            return is_mirror(l.left, r.right) and is_mirror(l.right, r.left)

        return is_mirror(root.left, root.right)


if __name__ == '__main__':
    root = create_tree_using_bfs([1, 2, 2, 3, 4, 4, 3])
    print(Solution().isSymmetric(root))


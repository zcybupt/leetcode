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

def dfs(root: TreeNode) -> None:
    if not root: return
    print(root.val, end = ' ')
    dfs(root.left)
    dfs(root.right)

def bfs(root: TreeNode) -> None:
    queue = [root]
    while queue:
        cur_node = queue[0]
        queue = queue[1:]
        print(cur_node.val, end = ' ')
        if cur_node.left: queue.append(cur_node.left)
        if cur_node.right: queue.append(cur_node.right)

if __name__ == '__main__':
    nums = [i for i in range(14)]
    root = create_tree_using_bfs(nums)
    bfs(root)

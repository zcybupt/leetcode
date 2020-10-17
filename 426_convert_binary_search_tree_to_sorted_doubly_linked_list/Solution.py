class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def createTreeUsingBFS(nums: list) -> Node:
    node_list = []

    for i, num in enumerate(nums):
        node_list.append(Node(num))
        if i == 0: continue
        if i % 2 == 1: node_list[(i - 1) // 2].left = node_list[i]
        else: node_list[(i - 1) // 2].right = node_list[i]

    return node_list[0]

def print_tree(root: Node):
    def dfs(root: Node):
        if not root: return
        dfs(root.left)
        print(root.val, end = ' ')
        dfs(root.right)
    dfs(root)
    print()


def print_list(head: Node):
    tmp = head
    while True:
        print(tmp.val, end = ' ')
        tmp = tmp.right
        if tmp == head: break


class Solution:
    def treeToDoublyList(self, root: Node) -> Node:
        if not root: return None

        def dfs(cur: Node):
            if not cur: return
            dfs(cur.left)
            if self.pre: self.pre.right = cur
            else: self.head = cur
            cur.left = self.pre
            self.pre = cur
            dfs(cur.right)

        self.pre = None
        dfs(root)
        self.head.left, self.pre.right = self.pre, self.head

        return self.head


if __name__ == '__main__':
    root = createTreeUsingBFS([10, 5, 15, 2, 6, 11, 16, 1, 3])
    print_tree(root)
    head = Solution().treeToDoublyList(root)
    print_list(head)


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def inorderTraversal(self, root: TreeNode) -> list:
        result = []
        if not root: return result

        def dfs(root: TreeNode):
            if not root: return
            dfs(root.left)
            result.append(root.val)
            dfs(root.right)

        dfs(root)

        return result


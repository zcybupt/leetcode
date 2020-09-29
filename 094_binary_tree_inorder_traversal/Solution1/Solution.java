import java.util.*;

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode() {};
    TreeNode(int val) { this.val = val; }
    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}

class Solution {
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        if (root == null) return result;

        dfs(root, result);

        return result;
    }

    public void dfs(TreeNode root, List<Integer> result) {
        if (root == null) return;

        dfs(root.left, result);
        result.add(root.val);
        dfs(root.right, result);
    }
}

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
    public static TreeNode createTreeUsingBFS(int[] nums) {
        TreeNode[] nodeArr = new TreeNode[nums.length];
        for (int i = 0; i < nums.length; i++) {
            nodeArr[i] = new TreeNode(nums[i]);
            if (i == 0) continue;
            if (i % 2 == 1) nodeArr[(i - 1) / 2].left = nodeArr[i];
            else nodeArr[(i - 1) / 2].right = nodeArr[i];
        }

        return nodeArr[0];
    }

    public boolean isSymmetric(TreeNode root) {
        if (root == null) return true;

        return isMirror(root.left, root.right);
    }

    public boolean isMirror(TreeNode l, TreeNode r) {
        if (l == null && r == null) return true;
        if (l == null || r == null || l.val != r.val) return false;

        return isMirror(l.left, r.right) && isMirror(l.right, r.left);
    }

    public static void main(String[] args) {
        TreeNode root = createTreeUsingBFS(new int[]{1, 2, 2, 3, 4, 4, 3});
        System.out.println(new Solution().isSymmetric(root));
    }
}

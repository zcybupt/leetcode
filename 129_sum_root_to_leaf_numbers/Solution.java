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

    public int sumNumbers(TreeNode root) {
        return dfs(root, 0);
    }

    public int dfs(TreeNode root, int sum) {
        if (root == null) return 0;
        if (root.left == null && root.right == null) return sum + root.val;
        sum += root.val;
        return dfs(root.left, 10 * sum) + dfs(root.right, 10 * sum);
    }

    public static void main(String[] args) {
        int[] nums = {4, 9, 0, 5, 1};
        TreeNode root = createTreeUsingBFS(nums);
        System.out.println(new Solution().sumNumbers(root));
    }
}


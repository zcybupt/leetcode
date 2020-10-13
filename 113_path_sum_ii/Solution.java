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

    public static void bfs(TreeNode root) {
        Deque<TreeNode> queue = new ArrayDeque<>();
        queue.addLast(root);
        while (!queue.isEmpty()) {
            TreeNode curNode = queue.pollFirst();
            System.out.print(curNode.val + " ");
            if (curNode.left != null) queue.addLast(curNode.left);
            if (curNode.right != null) queue.addLast(curNode.right);
        }
    }

    private List<List<Integer>> result;
    public List<List<Integer>> pathSum(TreeNode root, int sum) {
        result = new ArrayList<>();
        recur(root, new LinkedList<>(), sum);

        return result;
    }

    private void recur(TreeNode root, LinkedList<Integer> path, int sum) {
        if (root == null) return;
        path.add(root.val);
        sum -= root.val;
        if (sum == 0 && root.left == null && root.right == null)
            result.add(new ArrayList<>(path));
        recur(root.left, path, sum);
        recur(root.right, path, sum);
        path.removeLast();
    }

    public static void main(String[] args) {
        TreeNode root = createTreeUsingBFS(new int[]{5, 4, 8, 11, 13, 4, 7, 2, 5, 1});
        List<List<Integer>> result = new Solution().pathSum(root, 22);
        for (List<Integer> tmp : result)
            System.out.println(tmp);
    }
}

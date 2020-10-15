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

    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> result = new ArrayList<>();
        if (root == null) return result;

        Deque<TreeNode> queue = new ArrayDeque<>();
        queue.add(root);
        boolean flag = true;
        while (!queue.isEmpty()) {
            LinkedList<Integer> tmpRes = new LinkedList<>();
            for (int i = queue.size(); i > 0; i--) {
                TreeNode curNode = queue.pollFirst();
                if (flag) tmpRes.add(curNode.val);
                else tmpRes.addFirst(curNode.val);
                if (curNode.left != null) queue.add(curNode.left);
                if (curNode.right != null) queue.add(curNode.right);
            }
            flag = !flag;
            result.add(tmpRes);
        }

        return result;
    }

    public static void main(String[] args) {
        TreeNode root = createTreeUsingBFS(new int[]{1, 2, 3, 4, 5, 6, 7, 8, 9});
        List<List<Integer>> result = new Solution().zigzagLevelOrder(root);
        for (List<Integer> level : result) {
            System.out.println(level);
        }
    }
}

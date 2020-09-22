import java.util.*;

/**
 * 使用 used 数组标记路径，速度更快
 */
class Solution {
    List<List<Integer>> results = new ArrayList<>();

    public List<List<Integer>> subsets(int[] nums) {
        for (int i = 0; i <= nums.length; i++) {
            dfs(nums, 0, i, new ArrayDeque<>());
        }

        return results;
    }

    public void dfs(int[] nums, int start, int k, Deque<Integer> path) {
        if (k == 0) {
            results.add(new ArrayList<>(path));
            return;
        }

        for (int i = start; i < nums.length; i++) {
            path.addLast(nums[i]);
            dfs(nums, i + 1, k - 1, path);
            path.removeLast();
        }
    }

    public static void main(String[] args) {
        int[] nums = {1, 2, 3};
        List<List<Integer>> results = new Solution().subsets(nums);

        for (List<Integer> result : results) {
            System.out.println(result);
        }
    }
}

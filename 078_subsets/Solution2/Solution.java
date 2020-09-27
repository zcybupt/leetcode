import java.util.*;

/**
 * 使用 used 数组标记路径，速度更快
 */
class Solution {
    List<List<Integer>> results = new ArrayList<>();

    public List<List<Integer>> subsets(int[] nums) {
        for (int i = 0; i <= nums.length; i++) {
            boolean[] used = new boolean[nums.length];
            dfs(nums, 0, i, used);
        }

        return results;
    }

    public void dfs(int[] nums, int start, int k, boolean[] used) {
        if (k == 0) {
            List<Integer> tmpRes = new ArrayList<>();
            for (int i = 0; i < nums.length; i++) {
                if (used[i]) tmpRes.add(nums[i]);
            }
            results.add(tmpRes);
            return;
        }

        for (int i = start; i < nums.length; i++) {
            used[i] = true;
            dfs(nums, i + 1, k - 1, used);
            used[i] = false;
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

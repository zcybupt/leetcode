import java.util.*;

class Solution {
    public List<List<Integer>> results = new ArrayList<>();

    public List<List<Integer>> subsetsWithDup(int[] nums) {
        // 先排序，确保重复元素相邻
        Arrays.sort(nums);
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
            if (i > 0 && !used[i - 1] && nums[i] == nums[i - 1]) continue;
            used[i] = true;
            dfs(nums, i + 1, k - 1, used);
            used[i] = false;
        }
    }

    public static void main(String[] args) {
        int[] nums = {4, 4, 1, 4};
        List<List<Integer>> results = new Solution().subsetsWithDup(nums);

        for (List<Integer> result : results) {
            System.out.println(result);
        }
    }
}
